# windows-bench
Windows Bench

# Setup Windows Active Directory:

`Install-WindowsFeature dns -IncludeManagementTools`
`Install-WindowsFeature -Name ad-domain-services -IncludeManagementTools`


`Add-DnsServerPrimaryZone -Name "win-bench.pri" -ZoneFile "pri.win-bench.dns" -DynamicUpdate "NonsecureAndSecure"`

`Add-DnsServerPrimaryZone -NetworkID 192.168.33.0/24 -ZoneFile "33.168.192.in-addr.arpa.dns"  -DynamicUpdate "NonsecureAndSecure"`

`Set-ItemProperty -Path HKLM:\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters -Name 'Domain' -Value 'win-bench.pri'`

`Set-ItemProperty -Path HKLM:\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters -Name 'NV Domain' -Value 'win-bench.pri'`

# If server is running under Vagrant, you can disable the first Network Card like this:
`Disable-NetAdapter -Name "Ethernet" -Confirm:$false`

`Install-ADDSForest -DomainName "win-bench.pri" -SafeModeAdministratorPassword (Convertto-SecureString -AsPlainText "Winbench90" -Force)`


# Build:

`go get github.com/masterzen/winrm`

`GOOS=windows CGO_ENABLED=0  go build -o DefaultPasswordHistoryCount.exe main.go`

# Run on Windows:

`> DefaultPasswordHistoryCount.exe`
