package customer

import (
	"context"
	"github.com/DerryRenaldy/learnFiber/entity"
)

type CRepository interface {
	SaveCustomer(ctx context.Context, req *entity.Customer) (int64, error)
	SaveMerchant(ctx context.Context, req *entity.Customer) (bool, error)
	SavePhoneNumber(ctx context.Context, req *entity.Customer) (bool, error)
	SaveEmail(ctx context.Context, req *entity.Customer) (bool, error)
}
