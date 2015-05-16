package main

import (
	"encoding/json"
	"fmt"
	"github.com/kaneshin/giita/giita/client"
	"github.com/kaneshin/giita/giita/request"
)

func newTeamCommand() command {
	cmd := command{
		usage: "team",
	}
	cmd.run = teamCommand
	return cmd
}

func teamCommand(cmd *command, args []string) error {
	req := request.NewGetTeamRequest()
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
