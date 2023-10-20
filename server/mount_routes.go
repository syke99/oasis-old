package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/mux"
	"net/http"
)

func mountRoutesChi(mx *chi.Mux, endpoint Endpoint) *chi.Mux {
	// create a sub-router for this endpoint
	rtr := chi.NewRouter()

	rtr.Route(endpoint.Route, func(r chi.Router) {
		for method, handler := range endpoint.Handlers {
			// create method-specific sub-router
			sub := chi.NewRouter()

			// if the handler uses middleware, add said
			// middleware to handler
			if handler.Middleware != nil || len(handler.Middleware) != 0 {
				mw := make([]func(http.Handler) http.Handler, len(handler.Middleware))
				for i := range handler.Middleware {
					mw[i] = func(h http.Handler) http.Handler {
						return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
							handler.Middleware[i].ServeHTTP(w, r)
						})
					}
				}

				sub.Use(mw...)
			}

			// add the handler to the specified method
			sub.Method(string(method), "", func() http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					oW := NewOasisWriter(w, handler.Island)

					r.WithContext(context.WithValue(r.Context(), "props", handler.Island.GetProps()))

					handler.HandlerFunc(oW, r)
				})
			}())

			// mount the method-specific sub-router
			// to the main sub-router
			rtr.Mount("", sub)
		}
	})

	// mount the sub-router at the specified route
	mx.Mount(endpoint.Route, rtr)

	return mx
}

func mountRoutesGorilla(mx *mux.Router, endpoint Endpoint) *mux.Router {
	handlersByMethod := func(w http.ResponseWriter, r *http.Request) {
		handlerWithMiddleware := endpoint.Handlers[HTTPMethod(r.Method)]

		h := handlerWithMiddleware.HandlerFunc

		mw := handlerWithMiddleware.Middleware

		if mw != nil ||
			len(mw) > 0 {

			// server through middleware
			for i := range mw {
				m := mw[i]

				m.ServeHTTP(w, r)
			}
		}

		// hit last HandlerFunc
		h(w, r)
	}

	mx.Handle(endpoint.Route, handlerFuncToHandler(handlersByMethod))

	return mx
}

func handlerFuncToHandler(handler http.HandlerFunc) http.Handler {
	return handler
}
