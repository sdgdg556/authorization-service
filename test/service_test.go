package test

import (
	"auth-service/model"
	"auth-service/service"
	"context"
	"testing"
)

var (
	svc   *service.Service
	token string
)

const username = "caohaoyu"
const password = "123"
const role1 = "super"
const role2 = "normal"

func TestMain(m *testing.M) {
	svc = service.InitService()
	m.Run()
}

func TestAll(t *testing.T) {
	t.Run("TestCreateUser", TestCreateUser)
	t.Run("TestCreateRole", TestCreateRole)
	t.Run("TestUserAddRole", TestUserAddRole)
	t.Run("TestAuthorization", TestAuthorization)
	t.Run("TestUserCheckRole", TestUserCheckRole)
	t.Run("TestUserAllRoles", TestUserAllRoles)
	t.Run("TestDeleteUser", TestDeleteUser)
	t.Run("TestDeleteRole", TestDeleteRole)
	t.Run("TestInvalidate", TestInvalidate)
}

func TestCreateUser(t *testing.T) {
	req := model.UserRequest{
		Name:     username,
		Password: password,
	}
	if err := svc.CreateUser(context.Background(), &req); err != nil {
		t.Fatalf("CreateUser() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	if err := svc.CreateUser(context.Background(), &req); err == nil {
		t.Fatalf("CreateUser() 执行错误 ，期望err=user already exist! 实际结果err=nil")
	}

	t.Logf("CreateUser() 执行成功")
}

func TestCreateRole(t *testing.T) {
	req := model.RoleRequest{
		Name: role1,
	}
	if err := svc.CreateRole(context.Background(), &req); err != nil {
		t.Fatalf("CreateRole() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	if err := svc.CreateRole(context.Background(), &req); err == nil {
		t.Fatalf("CreateRole() 执行错误 ，期望err=role already exist! 实际结果err=nil")
	}
	req.Name = role2
	if err := svc.CreateRole(context.Background(), &req); err != nil {
		t.Fatalf("CreateRole() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	t.Logf("CreateRole() 执行成功")
}

func TestUserAddRole(t *testing.T) {
	req := model.UserAddRoleRequest{
		UserName:     username,
		UserPassword: password,
		RoleName:     role1,
	}
	if err := svc.UserAddRole(context.Background(), &req); err != nil {
		t.Fatalf("UserAddRole() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	req.RoleName = role2
	if err := svc.UserAddRole(context.Background(), &req); err != nil {
		t.Fatalf("UserAddRole() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	req.UserName = "caohaoyu123"
	if err := svc.UserAddRole(context.Background(), &req); err == nil {
		t.Fatalf("UserAddRole() 执行错误 ，期望err=user is invalid! 实际结果err=nil")
	}
	req.UserName = username
	req.UserPassword = "1234"
	if err := svc.UserAddRole(context.Background(), &req); err == nil {
		t.Fatalf("UserAddRole() 执行错误 ，期望err=user is invalid! 实际结果err=nil")
	}
	req.UserPassword = password
	req.RoleName = "xxxx"
	if err := svc.UserAddRole(context.Background(), &req); err == nil {
		t.Fatalf("UserAddRole() 执行错误 ，期望err=Role doesn't exist! 实际结果err=nil")
	}
	t.Logf("UserAddRole() 执行成功")
}

func TestAuthorization(t *testing.T) {
	req := model.UserRequest{
		Name:     username,
		Password: password,
	}
	resp, err := svc.Authorization(context.Background(), &req)
	if err != nil {
		t.Fatalf("Authorization() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	token = resp.Token

	resp2, err := svc.Authorization(context.Background(), &req)
	if err != nil {
		t.Fatalf("Authorization() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	if token != resp2.Token {
		t.Fatalf("Authorization() 执行错误 ，期望两次获取相同token 实际结果token1=%s, token2=%s", token, resp2.Token)
	}

	req.Name = "caohaoyu123"
	_, err = svc.Authorization(context.Background(), &req)
	if err == nil {
		t.Fatalf("Authorization() 执行错误 ，期望err=user is invalid! 实际结果err=nil")
	}
	req.Name = username
	req.Password = "1234"
	_, err = svc.Authorization(context.Background(), &req)
	if err == nil {
		t.Fatalf("Authorization() 执行错误 ，期望err=user is invalid! 实际结果err=nil")
	}
	t.Logf("Authorization() 执行成功")
}

func TestUserCheckRole(t *testing.T) {
	req := model.UserCheckRoleRequest{
		Token:    token,
		RoleName: role1,
	}
	resp, err := svc.UserCheckRole(context.Background(), &req)
	if err != nil {
		t.Fatalf("UserCheckRole() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	if !resp.Result {
		t.Fatalf("UserCheckRole() 执行错误 ，期望result=true 实际结果result=%v", resp.Result)
	}
	req.RoleName = role2
	resp, err = svc.UserCheckRole(context.Background(), &req)
	if err != nil {
		t.Fatalf("UserCheckRole() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	if !resp.Result {
		t.Fatalf("UserCheckRole() 执行错误 ，期望result=true 实际结果result=%v", resp.Result)
	}

	req.Token = "123"
	resp, err = svc.UserCheckRole(context.Background(), &req)
	if err == nil {
		t.Fatalf("UserCheckRole() 执行错误 ，期望err=Authentication faild! 实际结果err=nil")
	}
	if resp.Result {
		t.Fatalf("UserCheckRole() 执行错误 ，期望result=false 实际结果result=%v", resp.Result)
	}
	req.Token = token
	req.RoleName = "xxxx"
	resp, err = svc.UserCheckRole(context.Background(), &req)
	if err != nil {
		t.Fatalf("UserCheckRole() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	if resp.Result {
		t.Fatalf("UserCheckRole() 执行错误 ，期望result=false 实际结果result=%v", resp.Result)
	}
	t.Logf("UserCheckRole() 执行成功")
}

func TestUserAllRoles(t *testing.T) {
	req := model.TokenRequest{
		Token: token,
	}
	resp, err := svc.UserAllRoles(context.Background(), &req)
	if err != nil {
		t.Fatalf("UserAllRoles() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	if len(resp.Roles) != 2 {
		t.Fatalf("UserAllRoles() 执行错误 ，期望role数量为2 实际结果roles数量为%d", len(resp.Roles))
	}
	for _, v := range resp.Roles {
		if v != role1 && v != role2 {
			t.Fatalf("UserAllRoles() 执行错误 ，期望role为super, normal 实际结果roles为%+v", resp.Roles)
			break
		}
	}
	req.Token = "123"
	_, err = svc.UserAllRoles(context.Background(), &req)
	if err == nil {
		t.Fatalf("UserAllRoles() 执行错误 ，期望err=Authentication faild! 实际结果err=nil")
	}
	t.Logf("UserAddRole() 执行成功")
}

func TestDeleteUser(t *testing.T) {
	req := model.UserRequest{
		Name:     username,
		Password: password,
	}
	if err := svc.DeleteUser(context.Background(), &req); err != nil {
		t.Fatalf("DeleteUser() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	if err := svc.DeleteUser(context.Background(), &req); err == nil {
		t.Fatalf("DeleteUser() 执行错误 ，期望err=user is invalid! 实际结果err=nil")
	}
	req2 := model.UserAddRoleRequest{
		UserName:     username,
		UserPassword: password,
		RoleName:     role1,
	}
	if err := svc.UserAddRole(context.Background(), &req2); err == nil {
		t.Fatalf("UserAddRole() 执行错误 ，删除用户之后添加用户角色期望err=user is invalid! 实际结果err=nil")
	}
	t.Logf("DeleteUser() 执行成功")
}

func TestDeleteRole(t *testing.T) {
	req := model.RoleRequest{
		Name: role1,
	}
	if err := svc.DeleteRole(context.Background(), &req); err != nil {
		t.Fatalf("DeleteRole() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	if err := svc.DeleteRole(context.Background(), &req); err == nil {
		t.Fatalf("DeleteRole() 执行错误 ，期望err=Role doesn't exist! 实际结果err=nil")
	}
	req2 := model.UserAddRoleRequest{
		UserName:     username,
		UserPassword: password,
		RoleName:     role1,
	}
	if err := svc.UserAddRole(context.Background(), &req2); err == nil {
		t.Fatalf("UserAddRole() 执行错误 ，删除角色之后添加用户角色期望err=Role doesn't exist! 实际结果err=nil")
	}
	req.Name = role2
	if err := svc.DeleteRole(context.Background(), &req); err != nil {
		t.Fatalf("DeleteRole() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	t.Logf("DeleteRole() 执行成功")
}

func TestInvalidate(t *testing.T) {
	req := model.TokenRequest{
		Token: token,
	}
	if err := svc.Invalidate(context.Background(), &req); err != nil {
		t.Fatalf("Invalidate() 执行错误 ，期望err=nil 实际结果err=%+v", err)
	}
	resp, err := svc.UserAllRoles(context.Background(), &req)
	if err == nil {
		t.Fatalf("UserAllRoles() 执行错误 ，令牌过期之后查询所有角色期望err=Authentication faild! 实际结果err=nil")
	}
	if len(resp.Roles) > 0 {
		t.Fatalf("UserAllRoles() 执行错误 ，令牌过期之后查询所有角色期望roles为空 实际结果role=%+v", resp.Roles)
	}
	t.Logf("Invalidate() 执行成功")
}
