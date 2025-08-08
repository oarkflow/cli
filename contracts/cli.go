package contracts

type Cli interface {
	Register(commands []Command)
	Unregister(command string)
	Call(command string) error
	CallAndExit(command string) error
	Run(args []string, existIfCli bool) error
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
