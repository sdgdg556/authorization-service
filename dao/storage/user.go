package storage

import (
	"auth-service/helper"
	"auth-service/model"
	"context"
	"errors"
	"sync"
)

// CreateUser 创建用户
func (d *Dao) CreateUser(ctx context.Context, user *model.User) (err error) {
	// 用户存在
	if _, ok := d.GetUser(ctx, user.Name); ok {
		err = errors.New("user already exist! ")
		return
	}
	// 存储用户
	d.Users.Store(user.Name, user.Password)
	return
}

// DeleteUser 删除用户
func (d *Dao) DeleteUser(ctx context.Context, user *model.User) (err error) {
	// 校验用户名密码
	if err = d.validateUser(ctx, user); err != nil {
		return
	}
	// 删除用户
	d.Users.Delete(user.Name)
	return
}

// GetUser 获取用户信息
func (d *Dao) GetUser(ctx context.Context, userName string) (user *model.User, ok bool) {
	var (
		storePassword interface{}
	)
	// 用户不存在
	if storePassword, ok = d.Users.Load(userName); !ok {
		return
	}
	user = &model.User{
		Name:     userName,
		Password: storePassword.(string),
	}
	return
}

// UserAddRole 添加角色
func (d *Dao) UserAddRole(ctx context.Context, user *model.User, role *model.Role) (err error) {
	var (
		storeRoles interface{}
		ok         bool
		existRole  bool
	)
	// 校验用户名密码
	if err = d.validateUser(ctx, user); err != nil {
		return
	}
	// 角色不存在
	if _, ok = d.GetRole(ctx, role.Name); !ok {
		err = errors.New("Role doesn't exist! ")
		return
	}
	// 当前用户下不存在角色则直接新增
	if storeRoles, ok = d.UserRoles.Load(user.Name); !ok {
		roles := &sync.Map{}
		roles.Store(role.Name, struct{}{})
		d.UserRoles.Store(user.Name, roles)
		return
	}
	// 当前用户下存在角色判断是否已有角色
	storeRoles.(*sync.Map).Range(func(key, value interface{}) bool {
		if key.(string) == role.Name {
			existRole = true
			return false
		}
		return true
	})
	// 是则直接返回
	if existRole {
		return
	}
	// 新增角色
	storeRoles.(*sync.Map).Store(role.Name, struct{}{})
	d.UserRoles.Store(user.Name, storeRoles)
	return
}

// UserAllRoles 获取用户全部角色
func (d *Dao) UserAllRoles(ctx context.Context, token string) (roles []string, err error) {
	var (
		user       *model.User
		storeRoles interface{}
		ok         bool
	)
	if user, err = d.Authentication(ctx, token); err != nil {
		return
	}
	if storeRoles, ok = d.UserRoles.Load(user.Name); !ok {
		return
	}
	storeRoles.(*sync.Map).Range(func(key, value interface{}) bool {
		if _, ok = d.GetRole(ctx, key.(string)); ok {
			roles = append(roles, key.(string))
		}
		return true
	})
	return
}

// UserCheckRole 检查用户角色
func (d *Dao) UserCheckRole(ctx context.Context, token string, role *model.Role) (result bool, err error) {
	var (
		user       *model.User
		storeRoles interface{}
		ok         bool
	)
	if user, err = d.Authentication(ctx, token); err != nil {
		return
	}
	if storeRoles, ok = d.UserRoles.Load(user.Name); !ok {
		return
	}
	storeRoles.(*sync.Map).Range(func(key, value interface{}) bool {
		if key.(string) != role.Name {
			return true
		}
		if _, ok = d.GetRole(ctx, key.(string)); ok {
			result = true
			return false
		}
		return true
	})
	return
}

// validateUser 校验用户名密码
func (d *Dao) validateUser(ctx context.Context, user *model.User) (err error) {
	var (
		storeUser *model.User
		ok        bool
	)
	// 用户不存在
	if storeUser, ok = d.GetUser(ctx, user.Name); !ok {
		err = errors.New("user is invalid! ")
		return
	}
	// 密码校验失败
	if helper.Md5Encrypt(user.Password) != storeUser.Password {
		err = errors.New("user is invalid! ")
		return
	}
	return
}
