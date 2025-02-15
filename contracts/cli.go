package contracts

type Cli interface {
	Register(commands []Command)
	Unregister(command string)
	Call(command string)
	CallAndExit(command string)
	Run(args []string, existIfCli bool)
}

type Command interface {
	Signature() string
	Description() string
	Extend() Extend
	Handle(ctx Context) error
}

type Extend struct {
	Category string
	Flags    []Flag
}

type Context interface {
	Argument(index int) string
	Arguments() []string
	Option(key string) string
}

type Flag struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    string
}
