package http

import (
	"auth-service/model"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type Context struct {
	context.Context

	Request *http.Request
	Writer  http.ResponseWriter
}

func (c *Context) ParseJsonParams(ctx context.Context, req model.Request) (err error) {
	err = json.NewDecoder(c.Request.Body).Decode(req)
	if err != nil {
		c.JsonResponse(http.StatusForbidden,
			model.RawResponse(http.StatusForbidden, "Error in request!", nil),
		)
		return
	}
	if ok := req.Validate(); !ok {
		err = errors.New("Invalid params! ")
		c.JsonResponse(http.StatusBadRequest,
			model.RawResponse(http.StatusBadRequest, "Invalid params!", nil),
		)
		return
	}
	return
}

func (c *Context) JsonResponse(status int, response interface{}) (err error) {
	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Writer.WriteHeader(status)
	c.Writer.Header().Set("Content-Type", "application/json")
	_, err = c.Writer.Write(resp)
	return
}
