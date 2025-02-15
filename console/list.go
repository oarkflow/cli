package console

import (
	"github.com/oarkflow/cli/contracts"
)

func NewListCommand(app contracts.Cli) *ListCommand {
	return &ListCommand{app: app}
}

type ListCommand struct {
	app contracts.Cli
}

// Signature The name and signature of the console command.
func (receiver *ListCommand) Signature() string {
	return "list"
}

// Description The console command description.
func (receiver *ListCommand) Description() string {
	return "List commands"
}

// Extend The console command extend.
func (receiver *ListCommand) Extend() contracts.Extend {
	return contracts.Extend{}
}

// Handle Execute the console command.
func (receiver *ListCommand) Handle(ctx contracts.Context) error {
	receiver.app.Call("--help")

	return nil
}
