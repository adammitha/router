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
	srv.routes()
	srv.log.Infoln("Listening on port 8080...")
	http.ListenAndServe(":8080", srv.mux)
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
