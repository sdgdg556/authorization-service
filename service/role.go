package service

import (
	"auth-service/model"
	"context"
)

func (s *Service) CreateRole(ctx context.Context, req *model.RoleRequest) (err error) {
	role := &model.Role{
		Name: req.Name,
	}
	err = s.dao.CreateRole(ctx, role)
	return
}

func (s *Service) DeleteRole(ctx context.Context, req *model.RoleRequest) (err error) {
	role := &model.Role{
		Name: req.Name,
	}
	err = s.dao.DeleteRole(ctx, role)
	return
}
