package server

import (
	"github.com/syke99/oasis/islands"
	"net/http"
)

// Endpoint ties a route
// to a map of Handlers
// so that you can map a
// HandlerWithMiddleware
// to a specific HTTPMethod
type Endpoint struct {
	Route    string
	Handlers map[HTTPMethod]HandlerWithMiddleware
}

// HandlerWithMiddleware ties a http.HandlerFunc
// and islands.Island together, along with any
// middleware you want this handler to be passed
// through
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
