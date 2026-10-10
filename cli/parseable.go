package cli

type ParseableCliArgument interface {
	ParseCli(arg string, source any) error
}
