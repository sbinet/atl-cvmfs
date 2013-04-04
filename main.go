package main

import (
	"os"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

var g_cmd *commander.Commander

func init() {
	g_cmd = &commander.Commander{
		Name: os.Args[0],
		Commands: []*commander.Command{
			acvmfs_make_cmd_pkg_create(),
		},
		Flag:       flag.NewFlagSet("acvmfs", flag.ExitOnError),
		Commanders: []*commander.Commander{},
	}
}

func main() {

	err := g_cmd.Flag.Parse(os.Args[1:])
	handle_err(err)

	args := g_cmd.Flag.Args()
	err = g_cmd.Run(args)
	handle_err(err)

	return
}

// EOF
