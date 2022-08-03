package service

import (
	"auth-service/model"
	"context"
)

// Authorization 授权
func (s *Service) Authorization(ctx context.Context, req *model.UserRequest) (resp *model.TokenResponse, err error) {
	var token string
	user := &model.User{
		Name:     req.Name,
		Password: req.Password,
	}
	expire := model.ParseTime(s.authConfig.TokenExpire)
	resp = &model.TokenResponse{}
	if token, err = s.dao.Authorization(ctx, user, expire); err != nil {
		return
	}
	resp.Token = token
	return
}

// Invalidate 令牌过期
func (s *Service) Invalidate(ctx context.Context, req *model.TokenRequest) (err error) {
	return s.dao.Invalidate(ctx, req.Token)
}
