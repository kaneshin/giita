package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const (
	qiitaHost  = "qiita.com"
	apiVersion = "v2"
)

type Requester interface {
	ToRequest() (*http.Request, error)
}

type Request struct {
	Method string
	URL    url.URL
	Data   map[string]interface{}
}

func NewRequest(method, team, p string) Request {
	req := Request{}
	req.Method = method
	req.URL.Scheme = "https"
	if len(team) > 0 {
		req.URL.Host = fmt.Sprintf("%s.%s", team, qiitaHost)
	} else {
		req.URL.Host = qiitaHost
	}
	req.URL.Path = path.Join(path.Join("/api", apiVersion), p)
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
