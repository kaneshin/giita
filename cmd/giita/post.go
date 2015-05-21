package main

import (
	"fmt"
	"github.com/kaneshin/giita/giita/client"
	"github.com/kaneshin/giita/giita/request"
	"io/ioutil"
	"strings"
)

func newPostCommand() command {
	cmd := command{
		usage: "post <filename> -tags=foo,bar -title=\"Hello world\"",
	}
	cmd.run = postCommand
	return cmd
}

func postCommand(cmd *command, args []string) error {
	if len(args) == 0 {
		return cmd.err()
	}
	filename := args[0]
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var title string
	var tags []map[string]interface{}
	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, "-tags=") {
			tag := strings.TrimLeft(arg, "-tags")
			tag = strings.TrimLeft(tag, "=")
			for _, el := range strings.Split(tag, ",") {
				t := map[string]interface{}{
					"name":     el,
					"versions": []interface{}{},
				}
				tags = append(tags, t)
			}
		}
		if strings.HasPrefix(arg, "-title=") {
			title = strings.TrimLeft(arg, "-title=")
		}
	}
	if len(title) == 0 {
		title = "No Title."
	}

	req := request.NewPostItemRequest(team, title, string(dat))
	if len(tags) > 0 {
		req.Data["tags"] = tags
	}
	cli := client.NewClient(token)
	body, err := cli.Dispatch(req)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
