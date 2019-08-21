# Vagrant 
Vagrant

# Run:

`vagrant up`

# Setup Windows Active Directory:

`Install-WindowsFeature dns -IncludeManagementTools`
`Install-WindowsFeature -Name ad-domain-services -IncludeManagementTools`

`Disable-NetAdapter -Name "Ethernet" -Confirm:$false`

`Remove-NetIPAddress –InterfaceAlias "Ethernet 2" –IPAddress 192.168.22.10 –PrefixLength 24`

`New-NetIPAddress –InterfaceAlias "Ethernet 2" –IPAddress 192.168.22.10 –PrefixLength 24 -DefaultGateway 192.168.22.1`

`Set-DnsClientServerAddress -InterfaceAlias "Ethernet 2" -ServerAddresses 192.168.22.10`

`Disable-NetAdapterBinding –InterfaceAlias “Ethernet 2” –ComponentID ms_tcpip6`

`Add-DnsServerPrimaryZone -Name "win-bench.oss" -ZoneFile "oss.win-bench.dns" -DynamicUpdate "NonsecureAndSecure"`

`Add-DnsServerPrimaryZone -NetworkID 192.168.22.0/24 -ZoneFile "22.168.192.in-addr.arpa.dns"  -DynamicUpdate "NonsecureAndSecure"`

`Set-ItemProperty -Path HKLM:\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters -Name 'Domain' -Value 'win-bench.oss'`

`Set-ItemProperty -Path HKLM:\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters -Name 'NV Domain' -Value 'win-bench.oss'`

`Install-ADDSForest -DomainName "win-bench.oss" -SafeModeAdministratorPassword (Convertto-SecureString -AsPlainText "Vagrant90" -Force)`


### Join Domain Domain Controller

`Disable-NetAdapter -Name "Ethernet" -Confirm:$false`

`Remove-NetIPAddress –InterfaceAlias "Ethernet 2" –IPAddress 192.168.22.11 –PrefixLength 24`

`New-NetIPAddress –InterfaceAlias "Ethernet 2" –IPAddress 192.168.22.11 –PrefixLength 24 -DefaultGateway 192.168.22.1`

`Set-DnsClientServerAddress -InterfaceAlias "Ethernet 2" -ServerAddresses 192.168.22.10`

`Disable-NetAdapterBinding –InterfaceAlias “Ethernet 2” –ComponentID ms_tcpip6`

`Install-ADDSDomainController -DomainName "win-bench.oss" -SafeModeAdministratorPassword (Convertto-SecureString -AsPlainText "Vagrant90" -Force) -Credential  (Get-Credential win-bench\administrator)`


### Join Domain Add Computer

`Disable-NetAdapter -Name "Ethernet" -Confirm:$false`

`Remove-NetIPAddress –InterfaceAlias "Ethernet 2" –IPAddress 192.168.22.21 –PrefixLength 24`

`New-NetIPAddress –InterfaceAlias "Ethernet 2" –IPAddress 192.168.22.21 –PrefixLength 24 -DefaultGateway 192.168.22.1`

`Set-DnsClientServerAddress -InterfaceAlias "Ethernet 2" -ServerAddresses 192.168.22.10`

`Disable-NetAdapterBinding –InterfaceAlias “Ethernet 2” –ComponentID ms_tcpip6`

`Add-Computer –DomainName "win-bench.oss" -restart -Credential (Get-Credential win-bench\administrator)`



