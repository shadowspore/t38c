package t38c

// Command struct
type Command struct {
	Name string
	Args []interface{}
}

// NewCommand ...
func NewCommand(name string, args ...interface{}) Command {
	return Command{
		Name: name,
		Args: args,
	}
}
