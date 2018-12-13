# acctanon

Acctanon is a program written in Go that can be used to anonymise PowerMTA accounting files.

This program can be used to convert accounting files before sharing with another party, such as [Postmastery](https://www.postmastery.com). It is available as precompiled binary for Windows, Linux, and macOS. If required, users can also build the binary from source using the [Go distribution](https://golang.org/doc/install).

The local part of recipient addresses is replaced with "xxxx". In addition the local part is masked when found in VERP sender addresses and DSN messages. For example, given the following (simplified) input:

    type,orig,rcpt,dsnDiag
    b,noreply-pien=provider.com@sender.com,pien@provider.com,550 <pien@provider.com> Recipient not found

the converted output will be:

    type,orig,rcpt,dsnDiag
    b,noreply-xxxx=provider.com@sender.com,xxxx@provider.com,550 <xxxx@provider.com> Recipient not found

The program reads from stdin and writes to stdout. It can be used as follows:

    acctanon < acct-2018-12-10-0000.csv > acct-anon-2018-12-10-0000.csv

See the [releases](https://github.com/postmastery/acctanon/releases) tab for precompiled binaries. We recommend using /opt/pmta when installing on PowerMTA for Linux and \pmta\bin for Windows.

When building from source you need the [Go distribution](https://golang.org/doc/install). Then do:

    cd $GOPATH
    go get github.com/postmastery/acctanon
    cd src/github.com/postmastery/acctanon
    go install

Please [submit an issue](https://github.com/postmastery/acctanon/issues) or email the author if you have suggestions for improvement.
