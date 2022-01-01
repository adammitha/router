package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/adammitha/router"
	"github.com/sirupsen/logrus"
)

var port = flag.Int("port", 8080, "Port the server should listen on")

func main() {
	flag.Parse()
	srv := &Server{
		router.New(),
		logrus.New(),
		*port,
	}
	srv.mux.AddMiddleware(srv.logger)
	srv.routes()
	srv.log.Infof("Listening on port %d...", srv.port)
	http.ListenAndServe(fmt.Sprintf(":%d", srv.port), srv.mux)
}

type Server struct {
	mux  *router.Router
	log  *logrus.Logger
	port int
}

func (s *Server) routes() {
	s.mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})
}

func (s *Server) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.log.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL,
		}).Info()
		next.ServeHTTP(w, r)
	})
}
