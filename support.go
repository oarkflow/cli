package cli

import (
	"os"
	"slices"
)

var (
	version = "dev"
	cmd     = "cli"
)

func SetCommand(command string) {
	cmd = command
}

func SetVersion(ver string) {
	version = ver
}

func IsCli() bool {
	return slices.Contains(os.Args, cmd)
}
