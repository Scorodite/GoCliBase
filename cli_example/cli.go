package main

import (
	"fmt"

	"github.com/Scorodite/GoCliBase/cli"
)

type ExampleCli struct {
	*cli.CliEngine
}

func CreateExampleCli() *ExampleCli {
	res := &ExampleCli{
		CliEngine: cli.NewCliEngine(),
	}

	res.RegisterLibBase()
	res.RegisterLibStrings()
	res.RegisterLibVars()
	res.RegisterLibMath()
	res.RegisterLibFileSystem()
	res.RegisterLibConsole()

	res.RegisterMethod(res, "Sum", "Summarizes two digits", "a", "b")
	res.RegisterMethod(res, "PrintSeq", "Prints sequence elements", "Sequence")

	return res
}

func (c *ExampleCli) Sum(a, b float64) float64 {
	return a + b
}

func (c *ExampleCli) PrintSeq(s Sequence) Sequence {
	for i, a := range s.Items {
		fmt.Printf("%d| %s\n", i+1, a)
	}
	return s
}
