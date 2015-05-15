package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Requester interface {
	ToRequest() (*http.Request, error)
}

type Request struct {
	Method string
	URL    url.URL
	Data   map[string]interface{}
}

func NewRequest(method, team, path string) Request {
	req := Request{}
	req.Method = method
	req.URL.Scheme = "https"
	if len(team) > 0 {
		req.URL.Host = fmt.Sprintf("%s.qiita.com", team)
	} else {
		req.URL.Host = "qiita.com"
	}
	req.URL.Path = path
	req.Data = map[string]interface{}{}
	return req
}

func (r Request) ToRequest() (*http.Request, error) {
	data, err := json.Marshal(r.Data)
	if err != nil {
		return nil, err
	}
	return http.NewRequest(r.Method, r.URL.String(), bytes.NewBuffer(data))
}
