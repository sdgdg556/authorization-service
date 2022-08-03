package http

import (
	"auth-service/model"
	"net/http"
)

func CreateUserHandler(c *Context) {
	var (
		req     model.UserRequest
		code    = http.StatusOK
		message = "ok"
		data    interface{}
	)
	if err := c.ParseJsonParams(c, &req); err != nil {
		return
	}
	if err := s.CreateUser(c, &req); err != nil {
		code = http.StatusBadRequest
		message = err.Error()
	}
	c.JsonResponse(code, model.RawResponse(code, message, data))
}

func DeleteUserHandler(c *Context) {
	var (
		req     model.UserRequest
		code    = http.StatusOK
		message = "ok"
		data    interface{}
	)
	if err := c.ParseJsonParams(c, &req); err != nil {
		return
	}
	if err := s.DeleteUser(c, &req); err != nil {
		code = http.StatusBadRequest
		message = err.Error()
	}
	c.JsonResponse(code, model.RawResponse(code, message, data))
}

func UserAddRoleHandler(c *Context) {
	var (
		req     model.UserAddRoleRequest
		code    = http.StatusOK
		message = "ok"
		data    interface{}
	)
	if err := c.ParseJsonParams(c, &req); err != nil {
		return
	}
	if err := s.UserAddRole(c, &req); err != nil {
		code = http.StatusBadRequest
		message = err.Error()
	}
	c.JsonResponse(code, model.RawResponse(code, message, data))
}

func UserAllRolesHandler(c *Context) {
	var (
		req     model.TokenRequest
		code    = http.StatusOK
		message = "ok"
		data    interface{}
	)
	if err := c.ParseJsonParams(c, &req); err != nil {
		return
	}
	data, err := s.UserAllRoles(c, &req)
	if err != nil {
		code = http.StatusBadRequest
		message = err.Error()
	}
	c.JsonResponse(code, model.RawResponse(code, message, data))
}

func UserCheckRoleHandler(c *Context) {
	var (
		req     model.UserCheckRoleRequest
		code    = http.StatusOK
		message = "ok"
		data    interface{}
	)
	if err := c.ParseJsonParams(c, &req); err != nil {
		return
	}
	data, err := s.UserCheckRole(c, &req)
	if err != nil {
		code = http.StatusBadRequest
		message = err.Error()
	}
	c.JsonResponse(code, model.RawResponse(code, message, data))
}
