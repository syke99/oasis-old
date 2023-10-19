package server

import (
	"github.com/syke99/oasis/islands"
	"net/http"
)

type Endpoint struct {
	Route    string
	Handlers map[HTTPMethod]HandlerWithMiddleware
}

type HandlerWithMiddleware struct {
	HandlerFunc http.HandlerFunc
	Middleware  []http.HandlerFunc
	Island      islands.Island
}

type HTTPMethod string

const (
	MethodGet     = HTTPMethod(http.MethodGet)
	MethodPost    = HTTPMethod(http.MethodPost)
	MethodPut     = HTTPMethod(http.MethodPut)
	MethodDelete  = HTTPMethod(http.MethodDelete)
	MethodPatch   = HTTPMethod(http.MethodPatch)
	MethodConnect = HTTPMethod(http.MethodConnect)
	MethodTrace   = HTTPMethod(http.MethodTrace)
	MethodOptions = HTTPMethod(http.MethodOptions)
)

func NewEndpoint(route string, handlers map[HTTPMethod]HandlerWithMiddleware) Endpoint {
	return Endpoint{
		Route:    route,
		Handlers: handlers,
	}
}
