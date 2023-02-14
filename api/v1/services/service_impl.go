package services

import (
	"context"
	"database/sql"
	"github.com/DerryRenaldy/learnFiber/entity"
	"github.com/DerryRenaldy/learnFiber/forms"
	"github.com/gofiber/fiber/v2"
)

// GetCustomer Gets customer details by using phone number and merchant code
func (s service) GetCustomer(ctx context.Context, req forms.GetRequest) (*entity.Customer, error) {
	var resCusObj *entity.Customer
	if err := req.Validate(); err != nil {
		return nil, fiber.ErrInternalServerError
	}

	customerObj := &entity.Customer{
		MerchantCode: req.MerchantCode,
		PhoneNumber:  req.PhoneNumber,
	}

	resCusObj, err := s.customerRepo.GetByPhoneNumber(ctx, customerObj)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fiber.ErrInternalServerError
	}
	customerObj.ID = resCusObj.ID
	customerObj.Code = resCusObj.Code
	customerObj.MerchantCode = resCusObj.MerchantCode
	customerObj.PhoneNumber = resCusObj.PhoneNumber
	customerObj.Status = resCusObj.Status

	return customerObj, nil
}
