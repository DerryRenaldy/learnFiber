package services

import (
	"context"
	"github.com/DerryRenaldy/learnFiber/api/v1/repository/customer"
	"github.com/DerryRenaldy/learnFiber/entity"
	"github.com/DerryRenaldy/learnFiber/forms"
)

type service struct {
	customerRepo customer.CRepository
}

func New(customerRepo customer.CRepository) *service {
	obj := &service{
		customerRepo: customerRepo,
	}
	return obj
}

type IService interface {
	GetCustomer(ctx context.Context, req forms.GetRequest) (*entity.Customer, error)
}
