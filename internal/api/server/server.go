package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kranthi-reddy-gavireddy/internal/api/handler"
	"github.com/kranthi-reddy-gavireddy/internal/api/middlewares"
	"github.com/kranthi-reddy-gavireddy/internal/api/routes"
)

type Server struct {
	Handler handler.IHandler
	Router  *mux.Router
	Addr    string
}

func New(handler handler.IHandler) *Server {
	r := routes.NewRouter()
	routes.Register(r, handler)
	return &Server{
		Handler: handler,
		Router:  r,
		Addr:    ":8080",
	}
}

func (s *Server) AddMiddlewares() {
	s.Router.Use(middlewares.RequestLogger)
}

func (s *Server) Run() error {

	log.Printf("Server is running on %s", s.Addr)
	s.AddMiddlewares()
	return http.ListenAndServe(s.Addr, s.Router)
}
