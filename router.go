package router

import "net/http"

// Router is an HTTP multiplexer that allows you to register handlers
// for a specific URL pattern and HTTP method
type Router struct {
	mux        http.ServeMux
	middleware []func(http.Handler) http.Handler
}

// New instantiates a new Router instance
func New() *Router {
	return &Router{
		*http.NewServeMux(),
		make([]func(http.Handler) http.Handler, 0),
	}
}

// AddMiddleware registers a new middleware function that gets executed
// before every request.
func (r *Router) AddMiddleware(m func(http.Handler) http.Handler) {
	r.middleware = append(r.middleware, m)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
