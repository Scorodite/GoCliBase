package cli

type varsStdLib struct {
	Variables map[string]string
}

func (c *CliEngine) RegisterLibVars() {
	lib := &varsStdLib{
		Variables: map[string]string{},
	}

	c.RegisterMethod(lib, "Set", "Sets variable", "Name", "Value")
	c.RegisterMethod(lib, "Get", "Gets variable", "Name")
}

func (c *varsStdLib) Set(key, value string) string {
	c.Variables[key] = value
	return value
}

func (c *varsStdLib) Get(key string) string {
	return c.Variables[key]
}
