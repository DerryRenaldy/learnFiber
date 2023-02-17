package services

import (
	"database/sql"
	"github.com/DerryRenaldy/learnFiber/entity"
	"github.com/DerryRenaldy/learnFiber/forms"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

// GetCustomer Gets customer details by using phone number and merchant code
func (s service) GetCustomer(ctx *fasthttp.RequestCtx, req forms.GetRequest) (*entity.Customer, error) {
	functionName := "service.GetCustomer"
	var resCusObj *entity.Customer
	if err := req.Validate(); err != nil {
		s.l.Errorf("[%s - req.Validate] : %s", functionName, err)
		return nil, fiber.ErrInternalServerError
	}

	customerObj := &entity.Customer{
		MerchantCode: req.MerchantCode,
		PhoneNumber:  req.PhoneNumber,
	}

	resCusObj, err := s.customerRepo.GetByPhoneNumber(ctx, customerObj)
	if err == sql.ErrNoRows {
		s.l.Errorf("[%s - s.customerRepo.GetByPhoneNumber(NoRow)] : %s", functionName, err)
		return nil, nil
	} else if err != nil {
		s.l.Errorf("[%s - s.customerRepo.GetByPhoneNumber)] : %s", functionName, err)
		return nil, fiber.ErrInternalServerError
	}

	data, err := s.redisCache.GetByCode(ctx, resCusObj.ID)
	if err != nil {
		s.l.Errorf("[%s - s.redisCache] : %s", functionName, err)
	}

	s.l.Infof("[data from cache] : %s", data)

	customerObj.ID = resCusObj.ID
	customerObj.Code = resCusObj.Code
	customerObj.MerchantCode = resCusObj.MerchantCode
	customerObj.PhoneNumber = resCusObj.PhoneNumber
	customerObj.Email = resCusObj.Email
	customerObj.Status = resCusObj.Status

	return customerObj, nil
}
