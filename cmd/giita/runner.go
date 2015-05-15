package main

import (
	"flag"
	"fmt"
	"os"
)

var exitCode = 0

func main() {
	runnerMain()
	os.Exit(exitCode)
}

type runner struct {
	commands []command
}

func newRunner() runner {
	runner := runner{
		commands: []command{
			newPostCommand(),
		},
	}
	return runner
}

func (r *runner) finalize() {
}

func runnerMain() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		usage()
		exitCode = 1
		return
	}

	name := flag.Arg(0)

	if name == "--help" || name == "-h" {
		usage()
		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(stderr, "[ERROR]", r)
		}
	}()

	runner := newRunner()
	defer runner.finalize()

	for _, cmd := range runner.commands {
		if cmd.name() == name {
			if err := cmd.run(&cmd, flag.Args()[1:]); err != nil {
				exitCode = 1
				panic(err)
			}
			return
		}
	}
	exitCode = 1
	panic(fmt.Sprintf("Unknown command %s", name))
}
