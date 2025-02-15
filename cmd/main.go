package main

import (
	"os"

	"github.com/oarkflow/cli"
	"github.com/oarkflow/cli/console"
	"github.com/oarkflow/cli/contracts"
)

func main() {
	app := cli.New()
	client := app.Instance.Client()
	client.Register([]contracts.Command{
		console.NewListCommand(client),
	})
	client.Run(os.Args, true)
}
