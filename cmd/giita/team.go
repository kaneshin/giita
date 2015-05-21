package main

import (
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
	return result(body, &data)
}
