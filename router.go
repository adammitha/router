package router

import "net/http"

type Router struct {
	mux        http.ServeMux
	middleware []func(http.Handler) http.Handler
}

func New() *Router {
	return &Router{
		*http.NewServeMux(),
		make([]func(http.Handler) http.Handler, 0),
	}
}

func (r *Router) AddMiddleware(m func(http.Handler) http.Handler) {
	r.middleware = append(r.middleware, m)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
