package router

import "net/http"

func (r *Router) register(method, pattern string, handler http.Handler) {
	for _, m := range r.middleware {
		handler = m(handler)
	}
	r.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// Get registers a handler for a GET request for the provided pattern
func (r *Router) Get(pattern string, handler http.HandlerFunc) {
	r.register("GET", pattern, handler)
}

// Post registers a handler for a POST request for the provided pattern
func (r *Router) Post(pattern string, handler http.HandlerFunc) {
	r.register("POST", pattern, handler)
}

// Put registers a handler for a PUT request for the provided pattern
func (r *Router) Put(pattern string, handler http.HandlerFunc) {
	r.register("PUT", pattern, handler)
}

// Delete registers a handler for a DELETE request for the provided pattern
func (r *Router) Delete(pattern string, handler http.HandlerFunc) {
	r.register("DELETE", pattern, handler)
}
