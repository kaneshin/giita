package request

type PostItemRequest struct {
	Request
}

func NewPostItemRequest(team, title, body string) PostItemRequest {
	req := PostItemRequest{}
	req.Request = NewRequest("POST", team, "/api/v2/items")
	req.Data = map[string]interface{}{
		"body":      body,
		"coediting": false,
		"gist":      false,
		"private":   false,
		"tags":      []map[string]interface{}{},
		"title":     title,
		"tweet":     false,
	}
	return req
}
