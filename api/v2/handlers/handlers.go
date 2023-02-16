package handlers

import (
	serviceV2 "github.com/DerryRenaldy/learnFiber/api/v2/services"
	"github.com/DerryRenaldy/logger/logger"
	"github.com/gofiber/fiber/v2"
)

type CustomersHandler struct {
	l       logger.ILogger
	Service serviceV2.IService
}

type CustomersHandlerInterface interface {
	CreateCustomerHandler(c *fiber.Ctx) error
}

func NewCustomerHttpHandler(l logger.ILogger, customerService serviceV2.IService) *CustomersHandler {
	return &CustomersHandler{
		l:       l,
		Service: customerService,
	}
}
