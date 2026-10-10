package cli

import (
	"bufio"
	"os"
	"strings"
)

type consoleStdLib struct {
	reader *bufio.Reader
}

func (c *CliEngine) RegisterLibConsole() {
	lib := &consoleStdLib{
		reader: bufio.NewReader(os.Stdin),
	}

	c.RegisterMethod(lib, "Con.Read", "Reads line from console")
	c.RegisterMethod(lib, "Con.Write", "Writes text to console", "Text")
	c.RegisterMethod(lib, "Con.WriteLn", "Writes text to console, adding a new line", "Text")
}

// Reads line from console
func (sh *consoleStdLib) Read() string {
	readen, _ := sh.reader.ReadString('\n')
	return strings.TrimRight(readen, "\n\r")
}

// Writes text to console
func (sh *consoleStdLib) Write(text string) {
	print(text)
}

// Writes text to console, adding a new line
func (sh *consoleStdLib) WriteLn(text string) {
	println(text)
}
