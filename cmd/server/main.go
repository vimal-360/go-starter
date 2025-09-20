package main

import (
	"go-workflow-rnd/internal"
	"go-workflow-rnd/internal/config"
	"go-workflow-rnd/internal/routes"
	"log"

	"github.com/samber/do/v2"
)

func main() {
	// Create injector with single package loading
	injector := do.New(internal.Packages, internal.HandlerPackages)
	defer func() {
		if err := injector.Shutdown(); err != nil {
			log.Printf("Error during injector shutdown: %v", err)
		}
	}()

	// Get dependencies from the container
	cfg := do.MustInvoke[*config.Config](injector)

	e := routes.SetupRoutes(injector)

	log.Printf("Server starting on port %s", cfg.ServerPort)
	log.Fatal(e.Start(":" + cfg.ServerPort))
}
