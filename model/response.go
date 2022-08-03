package model

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TokenResponse struct {
	Token string `json:"auth_token"`
}

type RolesResponse struct {
	Roles []string `json:"roles"`
}

type UserCheckRoleResponse struct {
	Result bool `json:"result"`
}

func RawResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
