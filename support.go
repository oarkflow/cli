package cli

import (
	"os"
	"slices"
)

var (
	Version = "1.2.3"
	CliCmd  = "cli"
)

func IsCli() bool {
	return slices.Contains(os.Args, CliCmd)
}
