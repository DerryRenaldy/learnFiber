package services

import (
	"context"
	"github.com/DerryRenaldy/learnFiber/api/v1/repository/customer"
	"github.com/DerryRenaldy/learnFiber/entity"
	"github.com/DerryRenaldy/learnFiber/forms"
	"github.com/DerryRenaldy/logger/logger"
)

type service struct {
	l            logger.ILogger
	customerRepo customer.CRepository
}

func New(customerRepo customer.CRepository, log logger.ILogger) *service {
	obj := &service{
		customerRepo: customerRepo,
		l:            log,
	}
	return obj
}

//go:generate mockgen -source=./service.go -destination ./service_mocks.go -package=services
type IService interface {
	GetCustomer(ctx context.Context, req forms.GetRequest) (*entity.Customer, error)
}
