linux:
	GOARCH=amd64 GOOS=linux go build -o "bin/acctanon (Linux)"

windows:
	GOARCH=amd64 GOOS=windows go build -o "bin/acctanon.exe (Windows)"

darwin:
	GOARCH=amd64 GOOS=darwin go build -o "bin/acctanon (macOS)"

clean:
	rm -r bin/*

all: linux windows darwin
