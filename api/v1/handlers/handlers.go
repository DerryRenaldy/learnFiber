package handlers

import (
	serviceV1 "github.com/DerryRenaldy/learnFiber/api/v1/services"
	"github.com/DerryRenaldy/logger/logger"
	"github.com/gofiber/fiber/v2"
)

type CustomersHandler struct {
	l       logger.ILogger
	Service serviceV1.IService
}

type CustomersHandlerInterface interface {
	GetCustomerHandler(c *fiber.Ctx) error
}

func NewCustomerHttpHandler(l logger.ILogger, customerService serviceV1.IService) *CustomersHandler {
	return &CustomersHandler{
		l:       l,
		Service: customerService,
	}
}
