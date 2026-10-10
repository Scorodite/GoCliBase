# GoCliBase

Base library for creation command line interface
with custom commands and nested commands support

## Features
 - Execution code from string
 - Execution code from interactive console
 - Fast and convinient command definition
    - Supports custom types
 - Built-in libraries
    - Console interaction
    - File system interaction
    - Math
    - Strings
    - Global variables

## Syntax example

```
Con.Write "A: "
Set A [Con.Read]
Con.Write "B: "
Set B [Con.Read]

If [Calc [Get A] == "" || [Get B] == ""] [
    Echo A or B were not provided
] [
    Echo A + B = [Calc [Get A] + [Get B]]
]
```

## Quick start

```go
package main

import (
	"fmt"

	"github.com/Scorodite/GoCliBase/cli"
)

func main() {
	CreateExampleCli().EnterInteractiveMode()
}

type ExampleCli struct {
	*cli.CliEngine
}

func CreateExampleCli() *ExampleCli {
	res := &ExampleCli{
		CliEngine: cli.NewCliEngine(),
	}

	res.RegisterMethod(
        res, // Method owner
        "Sum", // Method name
        "Summarizes two digits", // About command
        "A", "B", // Argument names
    )

	return res
}

func (c *ExampleCli) Sum(a, b float64) float64 {
	return a + b
}
```

### Interactive console
```
>> Sum 10 3
        13
```

## Method commands

Methods can receive next arguments:
`int`, `int8-64`, `uint8-64`, `float32-64`, `bool`, `string`

Also, they can receive arguments that implement next interface:
```go
type ParseableCliArgument interface {
	ParseCli(arg string, source any) error
}
```
`source` is method owner, that was passed in `RegisterMethod`

Example of `ParseableCliArgument` implementation can be found at [sequence.go](cli_example/sequence.go)

Methods can return no value, any type single value, error or (any type, error)

You can find example code of method commands at [cli.go](cli_example/cli.go)

More complex examples of method commands can be found there:
 - [libbase.go](cli/libbase.go)
 - [libconsole.go](cli/libconsole.go)
 - [libfiles.go](cli/libfiles.go)
 - [libmath.go](cli/libmath.go)
 - [libstrings.go](cli/libstrings.go)
 - [libvars.go](cli/libvars.go)

## Example
This code can be executed with 'cli_example' package

```
Echo Break reason: [Loop [
    Set A ""; Set B ""

    Con.Write "Value A: "
    Set A [Con.Read]
    If [Calc [Get A] == ""] [Break "No A"]

    Con.Write "Value B: "
    Set B [Con.Read]
    If [Calc [Get B] == ""] [Break "No B"]

    Set Res [Try [Calc [Get A] + [Get B]]]

    If [IsError [Get Res]] [
        Echo Equasion failed!
        Continue
    ]
    
    Echo A + B = [Get Res]
    Echo Round(A + B) = [Math.Round [Get Res]]
]]

Set I 0
While [Calc [Get I] < 10] [
    Set I [Calc [Get I] + 1]
    Echo I is [Get I]
]
```

### Console output
```
Value A: 1
Value B: 2
A + B = 3
Round(A + B) = 3
Value A: 
Break reason: No A
I is 1
I is 2
...
I is 10
```

## Special commands

'If' executes body if condition is true

`If [Condition] [TrueBody]`

It can also execute false body if conditon is false
in such case other body will not be executed

`If [Condition] [TrueBody] [FalseBody]`


'While' executes body while condition is true

`While [Condition] [Body]`


'Loop' executes body repeatedly without condition. It must be interrupted with 'Break' command

`Loop [Body]`


'Break' stops cyclic commands from execution

`Break`


It can return value from loop if required

`Break [Value]`


'Continue' makes cyclic commands to execute
body from the beginning

`Continue`


'Try' attempts to execute body and returns error with "error:" prefix if execution was not successful

`Try [Body]`


'Yield' returns provided value

`Yield [Value]`

it can be used to return constant value inside subcommand:
```
If [Calc [Try [Command; Yield OK]] == OK] [
    Echo Completed successfuly
]
```
