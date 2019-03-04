default: build

build: build/linux/amd64 build/windows/x86

build/linux/amd64:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o workdir/milionerzy

build/windows/x86:
	GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o workdir/milionerzy.exe

test:
	go test -v
