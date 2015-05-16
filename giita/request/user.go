package request

type GetUserRequest struct {
	Request
}

func NewGetUserRequest(team string) GetUserRequest {
	req := GetUserRequest{}
	req.Request = NewRequest("GET", team, "/api/v2/users")
	req.Data = map[string]interface{}{}
	return req
}
