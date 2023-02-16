package forms

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type CreateRequest struct {
	MerchantCode       string `json:"merchantCode" validate:"required"`
	PhoneNumber        string `json:"phoneNumber" validate:"required"`
	Email              string `json:"email"`
	MerchantCustomerId string `json:"merchantCustomerId"`
	TransactionsId     string `json:"transactionsId"`
	ReferenceNumber    string `json:"referenceNumber"`
}

func (s *CreateRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("[ValidateCreateBadgeRequest] Invalid request in field [%+v], tag [%+v], value [%+v]", err.StructNamespace(), err.Tag(), err.Param())
		}
	}
	return nil
}
