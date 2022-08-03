package http

import (
	"auth-service/model"
	"net/http"
)

func CreateRoleHandler(c *Context) {
	var (
		req     model.RoleRequest
		code    = http.StatusOK
		message = "ok"
		data    interface{}
	)
	if err := c.ParseJsonParams(c, &req); err != nil {
		return
	}
	if err := s.CreateRole(c, &req); err != nil {
		code = http.StatusBadRequest
		message = err.Error()
	}
	c.JsonResponse(code, model.RawResponse(code, message, data))
}

func DeleteRoleHandler(c *Context) {
	var (
		req     model.RoleRequest
		code    = http.StatusOK
		message = "ok"
		data    interface{}
	)
	if err := c.ParseJsonParams(c, &req); err != nil {
		return
	}
	if err := s.DeleteRole(c, &req); err != nil {
		code = http.StatusBadRequest
		message = err.Error()
	}
	c.JsonResponse(code, model.RawResponse(code, message, data))
}
