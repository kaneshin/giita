package main

import (
	"errors"
	"fmt"
	"strings"
)

type command struct {
	usage string
	run   func(*command, []string) error
}

func (c *command) name() string {
	if f := strings.Fields(c.usage); len(f) > 0 {
		return f[0]
	}
	return ""
}

func (c *command) err() error {
	str := fmt.Sprintf("While executing %s %s\nUsage: %s", programName, c.name(), c.usage)
	return errors.New(str)
}
