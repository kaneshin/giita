package request

type GetTeamRequest struct {
	Request
}

func NewGetTeamRequest() GetTeamRequest {
	req := GetTeamRequest{}
	req.Request = NewRequest("GET", "", "/teams")
	req.Data = map[string]interface{}{}
	return req
}
