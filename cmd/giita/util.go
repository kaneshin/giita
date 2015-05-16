package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var programName = os.Args[0]

var team string
var token string
var conf = ".giita"

func init() {
	filename := path.Join(os.Getenv("HOME"), conf)
	if !Exists(filename) {
		fmt.Fprintf(stderr, "%s doesn't exist.\n", filename)
		os.Exit(1)
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var data map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}
	team = data["team"].(string)
	token = data["token"].(string)
}

var (
	stdout = os.Stdout
	stderr = os.Stderr
	stdin  = os.Stdin
)

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
