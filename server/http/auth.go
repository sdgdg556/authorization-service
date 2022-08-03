package http

import (
	"auth-service/model"
	"net/http"
)

func AuthorizationHandler(c *Context) {
	var (
		req     model.UserRequest
		code    = http.StatusOK
		message = "ok"
		data    interface{}
	)
	if err := c.ParseJsonParams(c, &req); err != nil {
		return
	}
	data, err := s.Authorization(c, &req)
	if err != nil {
		code = http.StatusBadRequest
		message = err.Error()
	}
	c.JsonResponse(code, model.RawResponse(code, message, data))
}

func InvalidateHandler(c *Context) {
	var (
		req     model.TokenRequest
		code    = http.StatusOK
		message = "ok"
		data    interface{}
	)
	if err := c.ParseJsonParams(c, &req); err != nil {
		return
	}
	err := s.Invalidate(c, &req)
	if err != nil {
		code = http.StatusBadRequest
		message = err.Error()
	}
	c.JsonResponse(code, model.RawResponse(code, message, data))
}
