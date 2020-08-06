package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/raulinoneto/catvotes/internal/adapters/primary"
	"github.com/raulinoneto/catvotes/internal/adapters/secondary"
	"github.com/raulinoneto/catvotes/pkg/domains/votes"
)

func main() {
	repo := secondary.NewVotesAPI()
	service := votes.NewService(repo)
	primaryPort := primary.NewHttpPrimaryAdapter(service)
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/vote", primaryPort.HandleVote)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}