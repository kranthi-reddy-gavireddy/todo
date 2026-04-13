package app

import (
	"github.com/kranthi-reddy-gavireddy/internal/api/handler"
	"github.com/kranthi-reddy-gavireddy/internal/api/repository"
	"github.com/kranthi-reddy-gavireddy/internal/api/server"
	"github.com/kranthi-reddy-gavireddy/internal/api/service"
)

func Run() error {
	// Initialize the repository, service, and handler
	repo := repository.New()
	service := service.New(repo)
	handler := handler.New(service)

	// Create and run the server
	server := server.New(handler)
	return server.Run()
}
