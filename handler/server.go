package handler

import (
	"sawitpro/pkg"
	"sawitpro/repository"
)

type Server struct {
	Repository repository.RepositoryInterface
	JWT        pkg.JWTInterface
}

type NewServerOptions struct {
	Repository repository.RepositoryInterface
	JWT        pkg.JWTInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		Repository: opts.Repository,
		JWT:        opts.JWT,
	}
}
