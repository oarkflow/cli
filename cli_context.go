package cli

import (
	"github.com/urfave/cli/v2"
)

type Context struct {
	instance *cli.Context
}

func (r *Context) Argument(index int) string {
	return r.instance.Args().Get(index)
}

func (r *Context) Arguments() []string {
	return r.instance.Args().Slice()
}

func (r *Context) Option(key string) string {

	return r.instance.String(key)
}
