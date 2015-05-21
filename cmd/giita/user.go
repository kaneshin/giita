package main

import (
	"github.com/kaneshin/giita/giita/client"
	"github.com/kaneshin/giita/giita/request"
)

func newUserCommand() command {
	cmd := command{
		usage: "user --page 1 --limit 20",
	}
	cmd.run = userCommand
	return cmd
}

func userCommand(cmd *command, args []string) error {
	id, ok := getOptionAsString(args, "--id")
	if ok {
		return getUserByID(id)
	}
	page, ok := getOptionAsInt(args, "--page")
	if !ok {
		page = 1
	}
	limit, ok := getOptionAsInt(args, "--limit")
	if !ok {
		limit = 20
	}
	return getUsers((int)(page), (int)(limit))
}

func getUsers(page, limit int) error {
	req := request.NewUserRequestWithPageAndLimit(team, page, limit)
	cli := client.NewClient(token)
	body, err := cli.Dispatch(req)
	if err != nil {
		return err
	}
	var data []map[string]interface{}
	return result(body, &data)
}

func getUserByID(id string) error {
	req := request.NewUserRequestWithID(team, id)
	cli := client.NewClient(token)
	body, err := cli.Dispatch(req)
	if err != nil {
		return err
	}
	var data map[string]interface{}
	return result(body, &data)
}
