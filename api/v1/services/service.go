package services

import (
	"context"
	customerV1 "github.com/DerryRenaldy/learnFiber/api/v1/repository/customer"
	"github.com/DerryRenaldy/learnFiber/api/v2/repository/customer_cache"
	"github.com/DerryRenaldy/learnFiber/entity"
	"github.com/DerryRenaldy/learnFiber/forms"
	"github.com/DerryRenaldy/logger/logger"
)

type service struct {
	l            logger.ILogger
	customerRepo customerV1.CRepository
	redisCache   customer_cache.IRepository
}

func New(customerRepo customerV1.CRepository, log logger.ILogger, redis customer_cache.IRepository) *service {
	obj := &service{
		customerRepo: customerRepo,
		l:            log,
		redisCache:   redis,
	}
	return obj
}

//go:generate mockgen -source=./service.go -destination ./service_mocks.go -package=services
type IService interface {
	GetCustomer(ctx context.Context, req forms.GetRequest) (*entity.Customer, error)
}
