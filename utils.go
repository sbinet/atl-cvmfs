package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func path_exists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func handle_err(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "**error**: %v\n", err.Error())
		//panic(err.Error())
		os.Exit(1)
	}
}

func _tar_gz(targ, workdir string, extra_args []string) error {
	// FIXME: use archive/tar instead (once go-1.1 is out)
	{
		matches, err := filepath.Glob(filepath.Join(workdir, "*"))
		if err != nil {
			return err
		}
		for i, m := range matches {
			matches[i] = m[len(workdir)+1:]
		}
		args := []string{}
		if len(extra_args) > 0 {
			args = append(args, extra_args...)
		}
		args = append(args, "-zcf", targ)
		args = append(args, matches...)
		//fmt.Printf(">> %v\n", args)
		cmd := exec.Command("tar", args...)
		cmd.Dir = workdir
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
}

// EOF
