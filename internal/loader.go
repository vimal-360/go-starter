package internal

import (
	"go-workflow-rnd/internal/config"
	"go-workflow-rnd/internal/database"
	"go-workflow-rnd/internal/handlers"
	"go-workflow-rnd/internal/repository"
	"go-workflow-rnd/internal/service"

	"github.com/samber/do/v2"
)

var Packages = do.Package(
	do.Lazy(config.NewConfigService),
	do.Lazy(database.NewDatabaseService),
	do.Lazy(repository.NewUserRepositoryService),
	do.Lazy(service.NewUserServiceDI),
	do.Lazy(handlers.NewUserHandlerService),
)

var HandlerPackages = do.Package(
	do.Lazy(handlers.NewUserHandlerService),
)
