package storage

import (
	"auth-service/helper"
	"auth-service/model"
	"context"
	"errors"
	"time"
)

// Authorization 授权,发放令牌
func (d *Dao) Authorization(ctx context.Context, user *model.User, expire time.Duration) (token string, err error) {
	var exist bool
	// 校验用户名密码
	if err = d.validateUser(ctx, user); err != nil {
		err = errors.New("Authorization faild! ")
		return
	}
	// user已存在授权token，直接返回
	d.tokens.Range(func(key, value interface{}) bool {
		if user.Name == value.(string) {
			token = key.(string)
			exist = true
			return false
		}
		return true
	})
	if exist {
		return
	}
	// 发放2小时过期的token
	token = d.BindToken(expire, user.Name)
	return
}

func (d *Dao) Invalidate(ctx context.Context, token string) (err error) {
	// token是否存在
	if _, ok := d.tokens.Load(token); !ok {
		err = errors.New("invalid auth token! ")
		return
	}
	// 删除token
	d.tokens.Delete(token)
	return
}

// BindToken 绑定令牌
func (d *Dao) BindToken(expire time.Duration, userName string) (token string) {
	token = helper.RandomString(36)
	d.tokens.Store(token, userName)
	time.AfterFunc(expire, func() {
		d.tokens.Delete(token)
	})
	return
}

// Authentication 鉴权，根据令牌返回对应用户
func (d *Dao) Authentication(ctx context.Context, token string) (user *model.User, err error) {
	var (
		storeUsername interface{}
		ok            bool
	)
	if storeUsername, ok = d.tokens.Load(token); !ok {
		err = errors.New("Authentication faild! ")
		return
	}
	// 获取用户信息看是否存在
	if user, ok = d.GetUser(ctx, storeUsername.(string)); !ok {
		err = errors.New("User doesn't exist! ")
		return
	}
	return
}
