package services

import (
	"context"
	customerV2 "github.com/DerryRenaldy/learnFiber/api/v2/repository/customer"
	"github.com/DerryRenaldy/learnFiber/entity"
	"github.com/DerryRenaldy/learnFiber/forms"
	"github.com/DerryRenaldy/logger/logger"
)

type service struct {
	l            logger.ILogger
	customerRepo customerV2.CRepository
}

func New(customerRepo customerV2.CRepository, log logger.ILogger) *service {
	obj := &service{
		customerRepo: customerRepo,
		l:            log,
	}
	return obj
}

//go:generate mockgen -source=./service.go -destination ./service_mocks.go -package=services
type IService interface {
	Create(ctx context.Context, req forms.CreateRequest) (*entity.Customer, error)
}
