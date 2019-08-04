# windows-bench
Windows Bench

build:

`go get github.com/masterzen/winrm`

run:

`GOOS=windows CGO_ENABLED=0  go build -o DefaultPasswordHistoryCount.exe main.go`