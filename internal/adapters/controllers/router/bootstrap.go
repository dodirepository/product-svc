package router

import (
	"github.com/dodirepository/product-svc/internal/adapters/controllers/handlers"
	"github.com/dodirepository/product-svc/internal/usecases"
)

var healthCheckControler = handlers.Healthcek{}

var itemController = &handlers.ItemHandlersController{
	Usecase: usecases.ItemUsecaseHandler(),
}
