package router

import "net/http"

type Router struct {
	mux        http.ServeMux
	middleware []func(http.HandlerFunc) http.HandlerFunc
}

func New() *Router {
	return &Router{
		*http.NewServeMux(),
		make([]func(http.HandlerFunc) http.HandlerFunc, 0),
	}
}

func (r *Router) AddMiddleware(m func(http.HandlerFunc) http.HandlerFunc) {
	_ = append(r.middleware, m)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
