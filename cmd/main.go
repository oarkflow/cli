package main

import (
	"github.com/oarkflow/cli"
	"github.com/oarkflow/cli/console"
	"github.com/oarkflow/cli/contracts"
)

func main() {
	err := cli.Run("CARE", "v0.0.1", RegisterCommands)
	if err != nil {
		panic(err)
	}
}

func RegisterCommands(client contracts.Cli) []contracts.Command {
	return []contracts.Command{
		console.NewListCommand(client),
	}
}
