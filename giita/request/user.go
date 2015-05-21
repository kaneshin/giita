package request

import (
	"path"
	"strconv"
)

type UserRequest struct {
	Request
}

func NewUserRequestWithPageAndLimit(team string, page, limit int) UserRequest {
	req := UserRequest{}
	req.Request = NewRequest("GET", team, "/users")
	req.Data = map[string]interface{}{
		"page":     strconv.Itoa(page),
		"per_page": strconv.Itoa(limit),
	}
	return req
}

func NewUserRequestWithID(team, id string) UserRequest {
	req := UserRequest{}
	req.Request = NewRequest("GET", team, path.Join("/users", id))
	req.Data = map[string]interface{}{}
	return req
}
