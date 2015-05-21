package main

import (
	"github.com/kaneshin/giita/giita/client"
	"github.com/kaneshin/giita/giita/request"
)

func newItemCommand() command {
	cmd := command{
		usage: "item --page 1 --limit 20",
	}
	cmd.run = itemCommand
	return cmd
}

func itemCommand(cmd *command, args []string) error {
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
	return getItems((int)(page), (int)(limit))
}

func getItems(page, limit int) error {
	req := request.NewItemRequestWithPageAndLimit(team, page, limit)
	cli := client.NewClient(token)
	body, err := cli.Dispatch(req)
	if err != nil {
		return err
	}
	var data []map[string]interface{}
	return result(body, &data)
}
