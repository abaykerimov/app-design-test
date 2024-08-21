package internal

import (
	"applicationDesignTest/handlers"
	"applicationDesignTest/internal/repositories"
	"applicationDesignTest/internal/services"
)

type Container struct {
	repository services.Repository
	service    handlers.Service
}

func NewContainer() *Container {
	repo := repositories.NewOrderRepository()
	service := services.NewService(repo)

	return &Container{
		repository: repo,
		service:    service,
	}
}
