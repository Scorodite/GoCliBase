package cli

type CliCallback func(args []string) (string, error)

type CliCommand struct {
	name      string
	arguments []string
	callback  CliCallback
	about     string
}

func (c CliCommand) GetName() string {
	return c.name
}

func (c CliCommand) GetArguments() []string {
	cp := make([]string, len(c.arguments))
	copy(cp, c.arguments)
	return cp
}

func (c CliCommand) GetCallback() CliCallback {
	return c.callback
}

func (c CliCommand) GetAbout() string {
	return c.about
}
