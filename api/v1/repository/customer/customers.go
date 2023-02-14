package customer

import (
	"context"
	"github.com/DerryRenaldy/learnFiber/entity"
)

type CRepository interface {
	GetByPhoneNumber(ctx context.Context, obj *entity.Customer) (*entity.Customer, error)
}
