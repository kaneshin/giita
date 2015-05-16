package request

type GetTeamRequest struct {
	Request
}

func NewGetTeamRequest() GetTeamRequest {
	req := GetTeamRequest{}
	req.Request = NewRequest("GET", "", "/api/v2/teams")
	req.Data = map[string]interface{}{}
	return req
}
