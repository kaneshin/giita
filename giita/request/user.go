package request

import (
	"path"
	"strconv"
)

type GetUserRequest struct {
	Request
}

func NewGetUserRequestWithPageAndLimit(team string, page, limit int) GetUserRequest {
	req := GetUserRequest{}
	req.Request = NewRequest("GET", team, "/users")
	req.Data = map[string]interface{}{
		"page":     strconv.Itoa(page),
		"per_page": strconv.Itoa(limit),
	}
	return req
}

func NewGetUserRequestWithID(team, id string) GetUserRequest {
	req := GetUserRequest{}
	req.Request = NewRequest("GET", team, path.Join("/users", id))
	req.Data = map[string]interface{}{}
	return req
}
