# Powershell Wrapper
Powershell Wrapper

# Build:

`go get github.com/robertojrojas/go-powershell`

`make`

# Run on Windows:

`PSWrapper.exe.exe "Get-WmiObject -Class Win32_Processor"`

`PSWrapper.exe.exe "Get-ADDefaultDomainPasswordPolicy -Identity win-bench.pri | Select -ExpandProperty PasswordHistoryCount"`
