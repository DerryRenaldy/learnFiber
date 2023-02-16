package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/DerryRenaldy/learnFiber/errors"
	"github.com/DerryRenaldy/learnFiber/forms"
	"github.com/DerryRenaldy/learnFiber/models"
	"github.com/DerryRenaldy/learnFiber/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"net/http"
	"regexp"
	"strings"
)

type NoCustomer struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Customer string `json:"customer"`
}

func (ch *CustomersHandler) CreateCustomerHandler(c *fiber.Ctx) error {
	functionName := "CustomersHandler.CreateCustomerHandler"
	ctx := c.Context()

	payload := &forms.CreateRequest{}
	if err := c.BodyParser(payload); err != nil {
		ch.l.Errorf("[%s - c.BodyParser] : %s", functionName, err)
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.CommonError{
			Code:    fiber.StatusInternalServerError,
			Message: fiber.ErrInternalServerError.Error(),
			Error:   err.Error(),
		})
	}

	if err := payload.Validate(); err != nil {
		ch.l.Errorf("[%s - payload.Validate] : %s", functionName, err)
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.CommonError{
			Code:    fiber.StatusInternalServerError,
			Message: fiber.ErrInternalServerError.Error(),
			Error:   err.Error(),
		})
	}

	validatedPayload, err := ch.CreateCustomerHelper(c.Context(), payload)
	if err != nil {
		ch.l.Errorf("[%s - CreateCustomerHelper] : %s", functionName, err)
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.CommonError{
			Code:    fiber.StatusInternalServerError,
			Message: fiber.ErrInternalServerError.Error(),
			Error:   err.Error(),
		})
	}

	customerObjFinal, err := ch.Service.Create(ctx, forms.CreateRequest{
		MerchantCode:       validatedPayload.MerchantCode,
		PhoneNumber:        validatedPayload.PhoneNumber,
		Email:              validatedPayload.Email,
		MerchantCustomerId: validatedPayload.MerchantCustomerId,
	})
	if err != nil {
		ch.l.Errorf("[%s - ch.Service.Create] : %s", functionName, err)
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.CommonError{
			Code:    fiber.StatusInternalServerError,
			Message: fiber.ErrInternalServerError.Error(),
			Error:   err.Error(),
		})
	}

	resultByte, _ := json.Marshal(customerObjFinal)
	ch.l.Debugf("[%s - Result Create Customer] : %s", functionName, string(resultByte))

	return c.Status(fiber.StatusCreated).JSON(&models.CreateCustomerResponse{
		Code:    fiber.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Customer: models.CreateCustomerDetail{
			Id:     customerObjFinal.Code,
			Status: customerObjFinal.Status,
		},
	})

}

func (ch *CustomersHandler) CreateCustomerHelper(ctx *fasthttp.RequestCtx, payload *forms.CreateRequest) (*forms.CreateRequest,
	error) {
	var err error

	payload.MerchantCode = strings.ToUpper(payload.MerchantCode)

	fmt.Println(payload)

	vTrxIdRefNumber := pkg.ValidateTxnIdRefNum(ctx, payload.TransactionsId, payload.ReferenceNumber)
	if vTrxIdRefNumber != nil {
		ch.l.Errorf("[Error Validation - pkg.ValidateTxnIdRefNum] : %s", vTrxIdRefNumber)
		return nil, fiber.ErrBadRequest
	}

	vMerchantCode := pkg.ValidateMerchantCode(ctx, payload.MerchantCode)
	if vMerchantCode != nil {
		ch.l.Errorf("[Error Validation - pkg.ValidateMerchantCode] : %s", vMerchantCode)
		return nil, fiber.ErrBadRequest
	}

	//Regex for Start with Digit and End with Digit
	var regexPhoneNumber = regexp.MustCompile("^[0-9]+$")
	if !regexPhoneNumber.MatchString(payload.PhoneNumber) {
		ch.l.Errorf("[Error Validation - no phone number] : no phone number in request")
		return nil, fiber.ErrBadRequest
	}

	payload.PhoneNumber, err = pkg.ValidatePhone(payload.PhoneNumber)
	if err != nil {
		ch.l.Errorf("[Error Validation - pkg.ValidatePhone] : %s", err)

		return nil, fiber.ErrBadRequest
	}

	if len(payload.Email) != 0 {
		if !pkg.ValidateEmail(payload.Email) {
			ch.l.Errorf("[Error Validation - pkg.ValidateEmail] : email is not valid")
			return nil, fiber.ErrBadRequest
		}
	}

	return payload, nil
}
