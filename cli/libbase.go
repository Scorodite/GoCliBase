package cli

import (
	"maps"
	"slices"
	"sort"
	"strings"
)

type cliStdLib struct {
	Engine *CliEngine
}

func (c *CliEngine) RegisterLibBase() {
	std := &cliStdLib{
		Engine: c,
	}

	c.RegisterMethod(std, "Help", "Prints all avariable commands")
	c.RegisterMethod(std, "IsError", "Checks if value has 'error:' prefix", "Value")

	c.Register("Calc", []string{"Args..."}, "Calculates math expressiona, can concatenate strings with '..' operator", func(args []string) (string, error) {
		return calc(args)
	})

	c.Register("Echo", []string{"Args..."}, "Prints arguments in console", func(args []string) (string, error) {
		text := strings.Join(args, " ")
		println(text)
		return text, nil
	})
}

func (c *cliStdLib) Help() {
	println("Avariable commands:")

	coms := c.Engine.GetCommands()
	keys := slices.Collect(maps.Keys(coms))
	sort.Slice(keys, func(i, j int) bool {
		return strings.Compare(keys[i], keys[j]) < 0
	})

	for _, key := range keys {
		com := coms[key]
		print("\t", com.GetName())

		for _, arg := range com.GetArguments() {
			print(" <", arg, ">")
		}

		if com.GetAbout() != "" {
			print("\t-\t", com.GetAbout())
		}

		println()
	}
}

func (c *cliStdLib) IsError(value string) bool {
	return strings.HasPrefix(value, "error:")
}
