# CLI

Small helper to build artisan-style CLIs on top of `urfave/cli/v2`, with a simple contract for commands and helpers for wiring, namespacing, and programmatic invocation.

## Installation

```bash
go get github.com/oarkflow/cli@latest
```

## Quick start

`cmd/main.go` already shows a minimal setup. The CLI namespacing keyword defaults to `cli`, so you run commands as `your-binary cli <command>`.

```go
package main

import (
    "github.com/oarkflow/cli"
    "github.com/oarkflow/cli/console"
    "github.com/oarkflow/cli/contracts"
)

func main() {
    // Optionally change the command namespace keyword (defaults to "cli").
    // cli.SetCommand("tool")

    // Run wires name, version, and your command registrations.
    if err := cli.Run("MyApp", "v1.0.0", register); err != nil {
        panic(err)
    }
}

func register(c contracts.Cli) []contracts.Command {
    return []contracts.Command{
        console.NewListCommand(c), // example command provided in this repo
    }
}
```

### Running

From the repo root:

```bash
go run ./cmd cli list        # show available commands via the sample list command
go run ./cmd cli --help      # global help
go run ./cmd cli --version   # prints the configured name and version
```

If you changed the namespace with `SetCommand("tool")`, call it as `go run ./cmd tool list`.

## Defining a command

Implement `contracts.Command`. Use `Extend` to add metadata like category and flags. `Handle` receives a `contracts.Context` for args and options.

```go
package console

import (
    "fmt"

    "github.com/oarkflow/cli/contracts"
)

type HelloCommand struct {
    app contracts.Cli
}

func NewHelloCommand(app contracts.Cli) *HelloCommand {
    return &HelloCommand{app: app}
}

func (h *HelloCommand) Signature() string   { return "hello" }
func (h *HelloCommand) Description() string { return "Print a greeting" }

func (h *HelloCommand) Extend() contracts.Extend {
    return contracts.Extend{
        Category: "examples",
        Flags: []contracts.Flag{
            {Name: "name", Aliases: []string{"n"}, Usage: "Name to greet", Value: "world"},
        },
    }
}

func (h *HelloCommand) Handle(ctx contracts.Context) error {
    fmt.Printf("Hello, %s!\n", ctx.Option("name"))
    return nil
}
```

Register it in your callback:

```go
func register(c contracts.Cli) []contracts.Command {
    return []contracts.Command{
        console.NewHelloCommand(c),
    }
}
```

Run it:

```bash
go run ./cmd cli hello --name Jane
```

## Helpful bits

- `cli.SetName(name)` and `cli.SetVersion(ver)` let you adjust metadata if you are not using `cli.Run`.
- `cli.SetCommand(command)` changes the namespace keyword (default `cli`).
- `contracts.Context` provides `Argument(i)`, `Arguments()`, and `Option(key)` for reading args and flags.
- `contracts.Cli` also offers `Call` and `CallAndExit` to invoke commands programmatically and `Unregister` to drop a command at runtime.
- `cli.IsCli()` quickly checks if the current process was invoked through the namespace keyword (e.g., `cli`).
