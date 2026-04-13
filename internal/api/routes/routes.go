package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kranthi-reddy-gavireddy/internal/api/handler"
)

type Path string

type Endpoint struct {
	Path
	Method  string
	Handler http.HandlerFunc
}

type Routes struct {
	Endpoints []Endpoint
}

func (r *Routes) AddRoute(path Path, method string, handler http.HandlerFunc) {
	r.Endpoints = append(r.Endpoints, Endpoint{
		Path:    path,
		Method:  method,
		Handler: handler,
	})
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	return r
}

func Register(router *mux.Router, h handler.IHandler) {
	router.HandleFunc("/todos", h.Create).Methods(http.MethodPost)
	router.HandleFunc("/todos/{id}", h.Update).Methods(http.MethodPut)
	router.HandleFunc("/todos/{id}", h.GetByID).Methods(http.MethodGet)
	router.HandleFunc("/todos", h.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/todos/{id}", h.Delete).Methods(http.MethodDelete)
}
