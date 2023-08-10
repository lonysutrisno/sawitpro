package main

import (
	"log"
	"os"
	"sawitpro/generated"
	"sawitpro/handler"
	"sawitpro/pkg"

	"sawitpro/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Panic occurred: %v", r)
		}
	}()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)

	// Start the server with error handling
	if err := e.Start(":1323"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	// dbDsn := "postgres://postgres:postgres@localhost:5432/database?sslmode=disable"
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})
	jwt := pkg.NewJWT()
	opts := handler.NewServerOptions{
		JWT:        jwt,
		Repository: repo,
	}
	return handler.NewServer(opts)
}
