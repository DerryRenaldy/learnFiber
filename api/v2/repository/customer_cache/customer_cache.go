package customer_cache

import (
	"github.com/DerryRenaldy/learnFiber/entity"
	"github.com/valyala/fasthttp"
)

type IRepository interface {
	// Set customer info into cache
	Set(ctx *fasthttp.RequestCtx, obj *entity.Customer) error
	// GetByCode get customer info from cache by customer code
	GetByCode(ctx *fasthttp.RequestCtx, customerCode int64) (*entity.Customer, error)
	// Remove customer info in cache by customer code
	Remove(ctx *fasthttp.RequestCtx, customerCode int64) error
	// GetList list of customer info in cache
	GetList(ctx *fasthttp.RequestCtx) ([]*entity.Customer, error)
}
