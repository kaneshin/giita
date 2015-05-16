package client

import (
	"github.com/kaneshin/giita/giita/request"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Token string
}

func NewClient(token string) Client {
	c := Client{}
	c.Token = token
	return c
}

func (c *Client) Dispatch(requester request.Requester) ([]byte, error) {
	req, err := requester.ToRequest()
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
