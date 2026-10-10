package main

import (
	"os"

	"github.com/Scorodite/GoCliBase/cli"
)

func main() {
	if len(os.Args) < 2 {
		CreateExampleCli().EnterInteractiveMode()
		return
	}

	code, e := os.ReadFile(os.Args[1])
	if e != nil {
		println("Failed to read file: ", e.Error())
		return
	}

	_, e = CreateExampleCli().Exec(string(code))
	if e != nil {
		println("Execution failed")
		cli.PrintCliError(string(code), e)
	}
}
