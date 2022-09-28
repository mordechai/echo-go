package server

import (
	"io"
	"log"
	"net/http"
	"strings"
)

// Server is a simple http server that echos GET requests

type Server struct {
	http.Server
}

// NewServer creates a new server
func NewServer() *Server {
	return &Server{}
}

// Listen starts the server on the given port
func (s *Server) Listen(port string) error {
	var err error
	http.HandleFunc("/", s.echo)

	log.Print("Server starting on: 127.0.0.1:", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		return err
	}
	defer s.Close()
	return nil
}

// echo echos the request
func (s *Server) echo(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, r.URL.Path)
	io.WriteString(w, strings.Split(r.URL.Path, "/")[1])
}
