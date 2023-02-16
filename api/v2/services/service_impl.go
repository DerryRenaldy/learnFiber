package services

import (
	"context"
	"github.com/DerryRenaldy/learnFiber/constant"
	"github.com/DerryRenaldy/learnFiber/entity"
	"github.com/DerryRenaldy/learnFiber/forms"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
)

func (s service) Create(ctx context.Context, req forms.CreateRequest) (*entity.Customer, error) {
	functionName := "service.Create"

	if err := req.Validate(); err != nil {
		s.l.Errorf("[%s - req.Validate] : %s", functionName, err)
		return nil, fiber.ErrInternalServerError
	}

	if strings.HasPrefix(req.PhoneNumber, "62") {
		req.PhoneNumber = req.PhoneNumber[2:]
	}

	phoneNumberAsInt, err := strconv.ParseInt(req.PhoneNumber, 10, 64)
	if err != nil {
		s.l.Errorf("[%s - strconv.ParseInt] : %s", functionName, err)
		return nil, fiber.ErrInternalServerError
	}

	phoneNumberEncode := strings.ToUpper(strconv.FormatInt(phoneNumberAsInt, 36))
	data := []string{req.MerchantCode, phoneNumberEncode}
	customerCode := strings.Join(data, "-")

	customerObj := &entity.Customer{
		Code:         customerCode,
		MerchantCode: req.MerchantCode,
		PhoneNumber:  "62" + req.PhoneNumber,
		Email:        req.Email,
		Status:       constant.CustomerStatusActive,
	}

	var customerId int64
	customerId, err = s.customerRepo.SaveCustomer(ctx, customerObj)
	if err != nil {
		if err.Error() == "Duplicate_Entry" {
			s.l.Errorf("[%s - s.customerRepo.SaveCustomer(Duplicate Customer)] : %s", functionName, err)
			return nil, fiber.ErrInternalServerError
		}
		s.l.Errorf("[%s - s.customerRepo.SaveCustomer] : %s", functionName, err)
		return nil, fiber.ErrInternalServerError
	}

	customerObj.ID = customerId
	merchantStatus, err := s.customerRepo.SaveMerchant(ctx, customerObj)
	if err != nil && !merchantStatus {
		s.l.Errorf("[%s - s.customerRepo.SaveMerchant] : %s", functionName, err)
		return nil, fiber.ErrInternalServerError
	}
	phoneStatus, err := s.customerRepo.SavePhoneNumber(ctx, customerObj)
	if err != nil && !phoneStatus {
		s.l.Errorf("[%s - s.customerRepo.SavePhoneNumber] : %s", functionName, err)
		return nil, fiber.ErrInternalServerError
	}
	emailStatus, err := s.customerRepo.SaveEmail(ctx, customerObj)
	if err != nil && !emailStatus {
		s.l.Errorf("[%s - s.customerRepo.SaveEmail] : %s", functionName, err)
		return nil, fiber.ErrInternalServerError
	}

	return customerObj, nil
}
