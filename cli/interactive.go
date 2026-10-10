package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (engine *CliEngine) EnterInteractiveMode() {
	reader := bufio.NewReader(os.Stdin)

	for {
		print(">> ")
		code, _ := reader.ReadString('\n')

		code = strings.TrimRight(code, "\r\n")
		for strings.HasSuffix(code, "\\") {
			print(".. ")
			extra, _ := reader.ReadString('\n')
			code += "\n"
			code += strings.TrimRight(extra, "\r\n")
		}

		res, e := engine.Exec(code)

		if e != nil {
			PrintCliError(code, e)
			continue
		}

		if res == "" {
			continue
		}

		fmt.Printf("\t%s\n\n", res)
	}
}
