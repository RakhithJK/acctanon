linux:
	GOARCH=amd64 GOOS=linux go build -o bin/linux_amd64/acctanon

windows:
	GOARCH=amd64 GOOS=windows go build -o bin/windows_amd64/acctanon.exe

darwin:
	GOARCH=amd64 GOOS=darwin go build -o bin/darwin_amd64/acctanon

all: linux windows darwin
