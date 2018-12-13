linux:
	GOARCH=amd64 GOOS=linux go build -o bin/linux_amd64/acctanon

windows:
	GOARCH=amd64 GOOS=windows go build -o bin/linux_amd64/windows_amd64/pdaimport.exe

darwin:
	GOARCH=amd64 GOOS=darwin go build -o bin/linux_amd64/darwin_amd64/pdaimport

all: linux windows darwin
