package cli

import (
	"github.com/oarkflow/cli/contracts"
)

type Application struct {
	instance contracts.Cli
}

func (app *Application) Init() contracts.Cli {
	app.instance = NewCli()
	return app.instance
}

func (app *Application) Client() contracts.Cli {
	return app.instance
}
