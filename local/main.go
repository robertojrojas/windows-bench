package main

import (
	"fmt"
	"log"
	"os"

	ps "github.com/robertojrojas/go-powershell"
	"github.com/robertojrojas/go-powershell/backend"
)

func main() {
	// choose a backend
	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	domain := "win-bench.pri"
	pshell := fmt.Sprintf("Get-ADDefaultDomainPasswordPolicy -Identity %s | Select -ExpandProperty PasswordHistoryCount", domain)

	if len(os.Args) > 1 {
		pshell = os.Args[1]
	}
	stdout, stderr, err := shell.Execute(pshell)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(stdout, stderr)
}
