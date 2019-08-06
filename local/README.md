# Powershell Wrapper
Powershell Wrapper

# Build:

`go get github.com/robertojrojas/go-powershell`

`GOOS=windows CGO_ENABLED=0  go build -o PSWrapper.exe main.go`

# Run on Windows:

`PSWrapper.exe.exe "Get-WmiObject -Class Win32_Processor"`

`PSWrapper.exe.exe "Get-ADDefaultDomainPasswordPolicy -Identity win-bench.pri | Select -ExpandProperty PasswordHistoryCount"`
