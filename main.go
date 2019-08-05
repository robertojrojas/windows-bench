package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/masterzen/winrm"
)

const powershellStr = `powershell.exe -Sta -NonInteractive -ExecutionPolicy RemoteSigned "%s"`

func main() {

	domain := "win-bench.pri"
	password := "vagrant"
	if len(os.Args) > 2 {
		domain = os.Args[1]
	}
	if len(os.Args) > 3 {
		password = os.Args[2]
	}
	host, _ := os.Hostname()

	client, err := setupClient(host, password)
	if err != nil {
		log.Fatal(err)
	}

	addPasswordPolicyQuery := fmt.Sprintf("Get-ADDefaultDomainPasswordPolicy -Identity %s | Select -ExpandProperty PasswordHistoryCount", domain)
	cmd := fmt.Sprintf(powershellStr, addPasswordPolicyQuery)
	var buffer bytes.Buffer
	var errBuffer bytes.Buffer
	fmt.Printf("Running host: %s %q\n", host, cmd)
	client.Run(cmd, &buffer, &errBuffer)
	fmt.Printf("PasswordHistoryCount on %s err: %q\n", buffer.String(), errBuffer.String())
}

func setupClient(host, password string) (*winrm.Client, error) {
	endpoint := winrm.NewEndpoint(host, 5985, false, false, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpoint, "Administrator", password)
	return client, err
}
