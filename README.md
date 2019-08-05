# windows-bench
Windows Bench

# Setup Windows Active Directory:

`Install-WindowsFeature dns -IncludeManagementTools`
`Install-WindowsFeature -Name ad-domain-services -IncludeManagementTools`

`Rename-computer –newname “dc1” -restart`

`Add-DnsServerPrimaryZone -Name "win-bench.pri" -ZoneFile "pri.win-bench.dns" -DynamicUpdate "NonsecureAndSecure"`

`Add-DnsServerPrimaryZone -NetworkID 192.168.33.0/24 -ZoneFile "33.168.192.in-addr.arpa.dns"  -DynamicUpdate "NonsecureAndSecure"`

`Set-ItemProperty -Path HKLM:\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters -Name 'Domain' -Value 'win-bench.pri'`

`Set-ItemProperty -Path HKLM:\SYSTEM\CurrentControlSet\Services\Tcpip\Parameters -Name 'NV Domain' -Value 'win-bench.pri'`

### If server is running under Vagrant, you can disable the first Network Card like this:
`Disable-NetAdapter -Name "Ethernet" -Confirm:$false`


`Install-ADDSForest -DomainName "win-bench.pri" -SafeModeAdministratorPassword (Convertto-SecureString -AsPlainText "Winbench90" -Force)`


# Preparing the remote Windows machine for Basic authentication
This project supports only basic authentication for local accounts (domain users are not supported). The remote windows system must be prepared for winrm:

_For a PowerShell script to do what is described below in one go, check [Richard Downer's blog](http://www.frontiertown.co.uk/2011/12/overthere-control-windows-from-java/)_

On the remote host, a PowerShell prompt, using the __Run as Administrator__ option and paste in the following lines:

		winrm quickconfig
		y
		winrm set winrm/config/service/Auth '@{Basic="true"}'
		winrm set winrm/config/service '@{AllowUnencrypted="true"}'
		winrm set winrm/config/winrs '@{MaxMemoryPerShellMB="1024"}'

__N.B.:__ The Windows Firewall needs to be running to run this command. See [Microsoft Knowledge Base article #2004640](http://support.microsoft.com/kb/2004640).

__N.B.:__ Do not disable Negotiate authentication as the `winrm` command itself uses this for internal authentication, and you risk getting a system where `winrm` doesn't work anymore.

__N.B.:__ The `MaxMemoryPerShellMB` option has no effects on some Windows 2008R2 systems because of a WinRM bug. Make sure to install the hotfix described [Microsoft Knowledge Base article #2842230](http://support.microsoft.com/kb/2842230) if you need to run commands that uses more than 150MB of memory.

For more information on WinRM, please refer to <a href="http://msdn.microsoft.com/en-us/library/windows/desktop/aa384426(v=vs.85).aspx">the online documentation at Microsoft's DevCenter</a>.




# Build:

`go get github.com/masterzen/winrm`

`GOOS=windows CGO_ENABLED=0  go build -o DefaultPasswordHistoryCount.exe main.go`

# Run on Windows:

`> DefaultPasswordHistoryCount.exe`
