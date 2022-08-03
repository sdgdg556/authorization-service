package http

import (
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type Handler struct {
	path    string
	handler HandlerFunc
}

type Route struct {
	handlers map[string][]*Handler
}

func InitRoute() *Route {
	return &Route{
		handlers: make(map[string][]*Handler),
	}
}

func (r *Route) AddRoute(method, path string, f HandlerFunc) {
	r.add(method, path, f)
}

func (r *Route) add(method, path string, f HandlerFunc) {
	h := &Handler{
		path:    strings.ToLower(path),
		handler: f,
	}
	r.handlers[strings.ToLower(method)] = append(r.handlers[strings.ToLower(method)], h)
}

func (r *Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := &Context{
		Writer:  w,
		Request: req,
	}
	url := strings.TrimLeft(req.URL.Path, "/")
	for _, handler := range r.handlers[strings.ToLower(req.Method)] {
		if handler.path == "/"+strings.ToLower(url) {
			switch strings.ToLower(req.Method) {
			case "get":
			case "post":
			}
			handler.ServeHTTP(c)
			return
		}
	}
	http.NotFound(w, req)
	return
}

func (h *Handler) ServeHTTP(c *Context) {
	h.handler(c)
}
