package handler

import (
	"coding-test-11-01-2025-dating-app/internal/repository"
)

type Server struct {
	Repository repository.RepositoryInterface
}

type NewServerOptions struct {
	Repository repository.RepositoryInterface
}

func NewServer(opts NewServerOptions) *Server {
	if opts.Repository == nil {
		panic("Repository cannot be nil")
	}

	return &Server{
		Repository: opts.Repository,
	}
}
