package main

import (
	"encoding/json"
	"fmt"
	"github.com/kaneshin/giita/giita/client"
	"github.com/kaneshin/giita/giita/request"
)

func newUserCommand() command {
	cmd := command{
		usage: "user",
	}
	cmd.run = userCommand
	return cmd
}

func userCommand(cmd *command, args []string) error {
	req := request.NewGetUserRequest(team)
	cli := client.NewClient(token)
	body, err := cli.Dispatch(req)
	if err != nil {
		return err
	}
	var data []map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(body))
	} else {
		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(body))
		} else {
			fmt.Println(string(b))
		}
	}
	return nil
}
