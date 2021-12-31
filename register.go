package router

import "net/http"

func (r *Router) register(method, pattern string, handler http.HandlerFunc) {
	for _, m := range r.middleware {
		handler = m(handler)
	}
	r.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	})
}

func (r *Router) Get(pattern string, handler http.HandlerFunc) {
	r.register("GET", pattern, handler)
}

func (r *Router) Post(pattern string, handler http.HandlerFunc) {
	r.register("POST", pattern, handler)
}

func (r *Router) Put(pattern string, handler http.HandlerFunc) {
	r.register("PUT", pattern, handler)
}

func (r *Router) Delete(pattern string, handler http.HandlerFunc) {
	r.register("DELETE", pattern, handler)
}
