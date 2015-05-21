package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
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

func parseArgs(args []string, name string) (string, bool) {
	for idx, arg := range args {
		if arg == name {
			if len(args) > idx+1 {
				return args[idx+1], true
			}
		}
	}
	return "", false
}

func getOptionAsString(args []string, name string) (string, bool) {
	return parseArgs(args, name)
}

func getOptionAsInt(args []string, name string) (int64, bool) {
	if opt, ok := parseArgs(args, name); ok {
		if val, err := strconv.ParseInt(opt, 10, 64); err == nil {
			return val, true
		}
	}
	return 0, false
}

func result(b []byte, d interface{}) error {
	err := json.Unmarshal(b, d)
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(b))
	} else {
		indb, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(b))
		} else {
			fmt.Println(string(indb))
		}
	}
	return err
}
