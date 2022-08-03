package storage

import (
	"auth-service/model"
	"context"
	"errors"
)

// CreateRole 创建角色
func (d *Dao) CreateRole(ctx context.Context, role *model.Role) (err error) {
	// 角色存在
	if _, ok := d.GetRole(ctx, role.Name); ok {
		err = errors.New("role already exist! ")
		return
	}
	// 存储角色
	d.Roles.Store(role.Name, struct{}{})
	return
}

// DeleteRole 删除角色
func (d *Dao) DeleteRole(ctx context.Context, role *model.Role) (err error) {
	// 角色不存在
	if _, ok := d.GetRole(ctx, role.Name); !ok {
		err = errors.New("Role doesn't exist! ")
		return
	}
	// 删除角色
	d.Roles.Delete(role.Name)
	return
}

// GetRole 获取角色信息
func (d *Dao) GetRole(ctx context.Context, roleName string) (role *model.Role, ok bool) {
	// 角色不存在
	if _, ok = d.Roles.Load(roleName); !ok {
		return
	}
	role = &model.Role{
		Name: roleName,
	}
	return
}
