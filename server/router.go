package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/syke99/oasis/islands"
	"golang.org/x/net/context"
	"net/http"
)

type Router struct {
	mux chi.Router
}

func NewRouter() *Router {
	return &Router{
		mux: chi.NewRouter(),
	}
}

func (r *Router) AddEndpoint(endpoint Endpoint) *Router {
	// create a sub-router for this endpoint
	rtr := chi.NewRouter()

	rtr.Route(endpoint.Route, func(r chi.Router) {
		for method, handler := range endpoint.Handlers {
			// create method-specific sub-router
			sub := chi.NewRouter()

			// if the handler uses middleware, add said
			// middleware to handler
			if handler.Middleware != nil || len(handler.Middleware) != 0 {
				for i := range handler.Middleware {
					sub.Use(func(h http.Handler) http.Handler {
						return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
							handler.Middleware[i].ServeHTTP(w, r)
						})
					})
				}
			}

			// add the handler to the specified method
			sub.Method(string(method), "", func() http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					oW := NewOasisWriter(w, handler.Island)

					r.WithContext(context.WithValue(r.Context(), "props", any(handler.Island).(*node).props.props))

					handler.HandlerFunc(oW, r)
				})
			}())

			// mount the method-specific sub-router
			// to the main sub-router
			rtr.Mount("", sub)
		}
	})

	// mount the sub-router at the specified route
	r.mux.Mount(endpoint.Route, rtr)

	return r
}

func PropsForRequest(r *http.Request) map[string]any {
	val := r.Context().Value("props")
	if val != nil {
		return val.(map[string]any)
	}

	return nil
}

func (r *Router) AddEndpoints(endpoints ...Endpoint) *Router {
	for i := range endpoints {
		r.AddEndpoint(endpoints[i])
	}
	return r
}

func UpgradeRouter(router chi.Router) *Router {
	return &Router{
		mux: router,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	r.mux.ServeHTTP(w, rq)
}

type OasisWriter interface {
	Write(p []byte) (n int, err error)
	Header() http.Header
	WriteHeader(statusCode int)
}

func NewOasisWriter(w http.ResponseWriter, island islands.Island) OasisWriter {
	return &oasisWriter{
		island: island,
		writer: w,
	}
}

type oasisWriter struct {
	island islands.Island
	writer http.ResponseWriter
}

func (o *oasisWriter) Header() http.Header {
	return o.writer.Header()
}

func (o *oasisWriter) WriteHeader(statusCode int) {
	o.writer.WriteHeader(statusCode)
}

func (o *oasisWriter) Write(p []byte) (n int, err error) {
	defer func() {
		if r := recover(); r != nil {
			n = 0
			err = fmt.Errorf("recovered from rendering island with err: %s", r.(error).Error())
		}
	}()

	payload := NewPayload()

	err = json.Unmarshal(p, payload.payload)
	if err != nil {
		return 0, err
	}

	o.island.Hydrate(payload.payload)

	return o.writer.Write([]byte(islands.MustRender(o.island)))
}

type OasisPayload struct {
	payload map[string]any
}

func NewPayload() *OasisPayload {
	return &OasisPayload{payload: make(map[string]any)}
}

func (p *OasisPayload) Set(key string, val any) {
	p.payload[key] = val
}

func (p *OasisPayload) Marshal() ([]byte, error) {
	return json.Marshal(p.payload)
}
