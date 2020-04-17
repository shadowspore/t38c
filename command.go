package t38c

// Command struct
type Command struct {
	Name string
	Args []string
}

// NewCommand return new command.
func NewCommand(name string, args ...string) Command {
	return Command{
		Name: name,
		Args: args,
	}
}

func (cmd Command) String() string {
	str := cmd.Name
	for _, arg := range cmd.Args {
		str += " " + arg
	}
	return str
}
