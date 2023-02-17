package services

import (
	customerV2 "github.com/DerryRenaldy/learnFiber/api/v2/repository/customer"
	"github.com/DerryRenaldy/learnFiber/api/v2/repository/customer_cache"
	"github.com/DerryRenaldy/learnFiber/entity"
	"github.com/DerryRenaldy/learnFiber/forms"
	"github.com/DerryRenaldy/logger/logger"
	"github.com/valyala/fasthttp"
)

type service struct {
	l            logger.ILogger
	customerRepo customerV2.CRepository
	redis        customer_cache.IRepository
}

func New(customerRepo customerV2.CRepository, log logger.ILogger, redisClient customer_cache.IRepository) *service {
	obj := &service{
		customerRepo: customerRepo,
		l:            log,
		redis:        redisClient,
	}
	return obj
}

//go:generate mockgen -source=./service.go -destination ./service_mocks.go -package=services
type IService interface {
	Create(ctx *fasthttp.RequestCtx, req forms.CreateRequest) (*entity.Customer, error)
}
