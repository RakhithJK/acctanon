package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

const mask = "xxxx"

func main() {

	// skip utf-8 byte order mark if present
	br := bufio.NewReader(os.Stdin)
	bom, err := br.Peek(3)
	if err != nil {
		return
	}
	if bom[0] == 0xEF && bom[1] == 0xBB && bom[2] == 0xBF {
		br.Discard(3)
	}
	r := csv.NewReader(br)
	r.ReuseRecord = true

	// read header line
	header, err := r.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if len(header) == 0 || header[0] != "type" {
		fmt.Fprintf(os.Stderr, "Expected header with \"type,...\"\n")
		os.Exit(1)
	}

	// find index of rcpt and dsnDiag
	var (
		origIndex    int
		rcptIndex    int
		dsnDiagIndex int
	)
	for i, field := range header {
		switch field {
		case "orig":
			origIndex = i
		case "rcpt":
			rcptIndex = i
		case "dsnDiag":
			dsnDiagIndex = i
		}
	}

	w := csv.NewWriter(os.Stdout)
	err = w.Write(header)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	for {
		// read next record
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		if rcptIndex != 0 {
			rcpt := record[rcptIndex]
			if rcpt != "" { // empty for tq records
				user, domain := splitEmail(rcpt)
				if user != "" {
					// mask local part in rcpt
					record[rcptIndex] = mask + "@" + domain
					// mask user in VERP address (e.g. jsmith-jdoe=yahoo.com@example.com)
					if origIndex != 0 {
						record[origIndex] = maskInVERP(record[origIndex], user)
					}
					// mask user in dsnDiag
					if dsnDiagIndex != 0 {
						record[dsnDiagIndex] = maskInDSN(record[dsnDiagIndex], user)
					}
				}
			}
		}

		err = w.Write(record)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}
	w.Flush()
}

// split email address in local and domain part
func splitEmail(email string) (user, domain string) {
	if i := strings.IndexRune(email, '@'); i != -1 {
		user = email[0:i]
		domain = strings.ToLower(email[i+1:])
	} else {
		domain = strings.ToLower(email) // no local part
	}
	return
}

// mask user in orig, e.g. jsmith-jdoe=yahoo.com@example.com
func maskInVERP(orig string, user string) string {
	if i := strings.Index(orig, user); i != -1 {
		if i+len(user) < len(orig) && orig[i+len(user)] == '=' {
			return orig[:i] + mask + orig[i+len(user):]
		}
	}
	return orig
}

// mask user in delivery status notification
func maskInDSN(dsn string, user string) string {
	// do quick find first, could match part of regular word (e.g. pien in recipient)
	if strings.Index(dsn, user) != -1 {
		// replace user when enclosed in word boundaries
		re, err := regexp.Compile(`\b` + regexp.QuoteMeta(user) + `\b`)
		if err != nil {
			panic(fmt.Sprintf("Error compiling regexp \\b%s\\b: %v", user, err))
		}
		return re.ReplaceAllString(dsn, mask)
	}
	return dsn
}
