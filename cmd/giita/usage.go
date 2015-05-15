package main

import (
	"fmt"
)

func usage() {
	fmt.Printf(`%s

  Usage:
  `, programName)

	for _, cmd := range newRunner().commands {
		fmt.Printf("  %s\n", cmd.usage)
	}

	fmt.Printf(`

  General Options:
    --verbose
`)
}
