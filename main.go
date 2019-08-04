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

	client, err := setupClient(password)
	if err != nil {
		log.Fatal(err)
	}

	addPasswordPolicyQuery := fmt.Sprintf("Get-ADDefaultDomainPasswordPolicy -Identity %s | Select -ExpandProperty PasswordHistoryCount", domain)
	cmd := fmt.Sprintf(powershellStr, addPasswordPolicyQuery)
	var buffer bytes.Buffer
	fmt.Printf("Running %q\n", cmd)
	client.Run(cmd, &buffer, os.Stderr)
	fmt.Printf("PasswordHistoryCount on %s\n", buffer.String())
}

func setupClient(password string) (*winrm.Client, error) {
	host, _ := os.Hostname()
	endpoint := winrm.NewEndpoint(host, 5985, false, false, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpoint, "Administrator", password)
	return client, err
}
