package request

import (
	"strconv"
)

type ItemRequest struct {
	Request
}

func NewItemRequestWithPageAndLimit(team string, page, limit int) ItemRequest {
	req := ItemRequest{}
	req.Request = NewRequest("GET", team, "/items")
	req.Data = map[string]interface{}{
		"page":     strconv.Itoa(page),
		"per_page": strconv.Itoa(limit),
	}
	return req
}
