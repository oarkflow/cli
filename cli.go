package cli

import (
	"os"
	"slices"
	"strings"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"

	"github.com/oarkflow/cli/contracts"
)

var cliName = "Boilerplate"

func SetName(name string) {
	cliName = name
}

type Callback func(client contracts.Cli) []contracts.Command

func Run(name, version string, callback Callback) error {
	SetName(name)
	SetVersion(version)
	app := New()
	client := app.Instance.Client()
	commands := callback(client)
	client.Register(commands)
	return client.Run(os.Args, true)
}

type Cli struct {
	instance *cli.App
}

func NewCli(name ...string) *Cli {
	if len(name) > 0 {
		cliName = name[0]
	}
	instance := cli.NewApp()
	instance.Name = cliName
	instance.Usage = version
	instance.UsageText = cmd + " [global options] command [options] [arguments...]"

	return &Cli{instance}
}

func (c *Cli) Register(commands []contracts.Command) {
	for _, item := range commands {
		item := item
		cliCommand := cli.Command{
			Name:  item.Signature(),
			Usage: item.Description(),
			Action: func(ctx *cli.Context) error {
				return item.Handle(&Context{ctx})
			},
		}

		cliCommand.Category = item.Extend().Category
		cliCommand.Flags = flagsToCliFlags(item.Extend().Flags)
		c.instance.Commands = append(c.instance.Commands, &cliCommand)
	}
}

func (c *Cli) Unregister(command string) {
	for idx, v := range c.instance.Commands {
		if v.Name == command {
			c.instance.Commands = append(c.instance.Commands[0:idx], c.instance.Commands[idx+1:]...)
			break
		}
	}
}

// Call Run an Artisan console command by name.
func (c *Cli) Call(command string) error {
	return c.Run(append([]string{os.Args[0], cmd}, strings.Split(command, " ")...), false)
}

// CallAndExit Run an Artisan console command by name and exit.
func (c *Cli) CallAndExit(command string) error {
	return c.Run(append([]string{os.Args[0], cmd}, strings.Split(command, " ")...), true)
}

// Run a command. Args come from os.Args.
func (c *Cli) Run(args []string, exitIfCli bool) error {
	if len(args) >= 2 {
		if index := slices.Index(args, cmd); index != -1 {
			cmdIndex := index + 1
			if len(args) == cmdIndex {
				args = append(args, "--help")
			}
			if args[cmdIndex] != "-V" && args[cmdIndex] != "--version" {
				cliArgs := append([]string{args[0]}, args[cmdIndex:]...)
				if err := c.instance.Run(cliArgs); err != nil {
					return err
				}
			}
			printResult(args[cmdIndex])
			if exitIfCli {
				os.Exit(0)
			}
		}
	}
	return nil
}

func flagsToCliFlags(flags []contracts.Flag) []cli.Flag {
	var cliFlags []cli.Flag
	for _, flag := range flags {
		cliFlags = append(cliFlags, &cli.StringFlag{
			Name:     flag.Name,
			Aliases:  flag.Aliases,
			Usage:    flag.Usage,
			Required: flag.Required,
			Value:    flag.Value,
		})
	}

	return cliFlags
}

func printResult(command string) {
	switch command {
	case "make:command":
		color.Greenln("Console command created successfully")
	case "-V", "--version":
		color.Greenln(cliName + " " + version)
	}
}
