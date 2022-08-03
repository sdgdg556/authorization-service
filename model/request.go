package model

type Request interface {
	Validate() bool
}

type UserRequest struct {
	Name     string `json:"user_name"`
	Password string `json:"user_password"`
}

type RoleRequest struct {
	Name string `json:"role_name"`
}

type UserAddRoleRequest struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	RoleName     string `json:"role_name"`
}

type UserCheckRoleRequest struct {
	Token    string `json:"auth_token"`
	RoleName string `json:"role_name"`
}

type TokenRequest struct {
	Token string `json:"auth_token"`
}

func (req *UserRequest) Validate() bool {
	if len(req.Name) == 0 || len(req.Password) == 0 {
		return false
	}
	return true
}

func (req *RoleRequest) Validate() bool {
	if len(req.Name) == 0 {
		return false
	}
	return true
}

func (req *UserAddRoleRequest) Validate() bool {
	if len(req.UserName) == 0 || len(req.UserPassword) == 0 || len(req.RoleName) == 0 {
		return false
	}
	return true
}

func (req *TokenRequest) Validate() bool {
	if len(req.Token) == 0 {
		return false
	}
	return true
}

func (req *UserCheckRoleRequest) Validate() bool {
	if len(req.Token) == 0 || len(req.RoleName) == 0 {
		return false
	}
	return true
}
