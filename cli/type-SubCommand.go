package cli

// SubCommand is a subcommand to a cli
type SubCommand struct {
	CLI
}

func newSubCommand(name string, desc string, fn commandFn, paramNames ...string) *SubCommand {
	var cmd SubCommand
	cmd.init(name, desc, fn, paramNames...)
	return &cmd
}
