# acctanon

Acctanon is a program written in Go that can be used to anonymise PowerMTA accounting files.

This program can be used to convert accounting files before sharing with another party, such as Postmastery. It is available in binary versions for Windows, Linux, and macOS. If required, users can also build the binary from source using the [Go distribution](https://golang.org/doc/install).

The local part of recipient addresses is replaced with "xxxx". In addition the local part is masked when found in VERP sender addresses and DSN messages. Given the following (simplified) input for example:

    type,orig,rcpt,dsnDiag
    b,noreply-pien=provider.com@sender.com,pien@provider.com,550 <pien@provider.com> Recipient not found

The converted output will be:

    type,orig,rcpt,dsnDiag
    b,noreply-xxxx=hotmail.com@sender.com,xxxx@hotmail.com,550 <xxxx@hotmail.com> Recipient not found

The program reads from stdin and writes to stdout. It can be used as follows:

    acctanon < acct-2018-12-10-0000.csv > acct-anon-2018-12-10-0000.csv


