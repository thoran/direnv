package main

import (
	"fmt"
	"os"
	"strconv"
)

// `direnv dump`
var CmdDump = &Cmd{
	Name:    "dump",
	Desc:    "Used to export the inner bash state at the end of execution",
	Args:    []string{"[SHELL]", "[FD]"},
	Private: true,
	Action: actionSimple(func(env Env, args []string) (err error) {
		target := "gzenv"
		out := os.Stdout

		if len(args) > 1 {
			target = args[1]
		}

		if len(args) > 2 {
			fd, err := strconv.ParseUint(args[2], 10, 32)
			if err != nil {
				return err
			}
			out = os.NewFile(uintptr(fd), "custom")
			if out == nil {
				return fmt.Errorf("fd %s is not valid", args[2])
			}
		}

		shell := DetectShell(target)
		if shell == nil {
			return fmt.Errorf("unknown target shell '%s'", target)
		}

		out.WriteString(shell.Dump(env))

		return
	}),
}
