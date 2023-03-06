package customer

import (
	"context"
	"github.com/DerryRenaldy/learnFiber/entity"
	"github.com/DerryRenaldy/learnFiber/forms"
	"github.com/DerryRenaldy/learnFiber/models"
)

type CRepository interface {
	SaveCustomer(ctx context.Context, req *entity.Customer) (int64, error)
	SaveMerchant(ctx context.Context, req *entity.Customer) (bool, error)
	SavePhoneNumber(ctx context.Context, req *entity.Customer) (bool, error)
	SaveEmail(ctx context.Context, req *entity.Customer) (bool, error)
	UpdateByCustomerId(ctx context.Context, obj forms.UpdateRequest) (*models.CustomerItemUpdate, error)
}
