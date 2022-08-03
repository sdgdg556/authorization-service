package service

import (
	"auth-service/model"
	"context"
)

func (s *Service) CreateUser(ctx context.Context, req *model.UserRequest) (err error) {
	user := &model.User{
		Name:     req.Name,
		Password: req.Password,
	}
	user.EncryptPassword()
	return s.dao.CreateUser(ctx, user)
}

func (s *Service) DeleteUser(ctx context.Context, req *model.UserRequest) (err error) {
	user := &model.User{
		Name:     req.Name,
		Password: req.Password,
	}
	return s.dao.DeleteUser(ctx, user)
}

func (s *Service) UserAddRole(ctx context.Context, req *model.UserAddRoleRequest) (err error) {
	user := &model.User{
		Name:     req.UserName,
		Password: req.UserPassword,
	}
	role := &model.Role{
		Name: req.RoleName,
	}
	return s.dao.UserAddRole(ctx, user, role)
}

func (s *Service) UserAllRoles(ctx context.Context, req *model.TokenRequest) (resp *model.RolesResponse, err error) {
	var (
		roles []string
	)
	resp = &model.RolesResponse{}
	if roles, err = s.dao.UserAllRoles(ctx, req.Token); err != nil {
		return
	}
	resp.Roles = roles
	return
}

func (s *Service) UserCheckRole(ctx context.Context, req *model.UserCheckRoleRequest) (resp *model.UserCheckRoleResponse, err error) {
	var (
		result bool
	)
	role := &model.Role{
		Name: req.RoleName,
	}
	resp = &model.UserCheckRoleResponse{}
	if result, err = s.dao.UserCheckRole(ctx, req.Token, role); err != nil {
		return
	}
	resp.Result = result
	return
}
