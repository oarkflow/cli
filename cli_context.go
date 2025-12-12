package cli

import (
	"github.com/urfave/cli/v3"
)

type Context struct {
	command *cli.Command
}

func (r *Context) Argument(index int) string {
	return r.command.Args().Get(index)
}

func (r *Context) Arguments() []string {
	return r.command.Args().Slice()
}

func (r *Context) Option(key string) string {
	return r.command.String(key)
}
