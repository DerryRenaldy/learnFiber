package handlers

import (
	"fmt"
	"github.com/DerryRenaldy/learnFiber/constant"
	"github.com/DerryRenaldy/learnFiber/errors"
	"github.com/DerryRenaldy/learnFiber/forms"
	"github.com/DerryRenaldy/learnFiber/models"
	"github.com/DerryRenaldy/learnFiber/pkg/tracer"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type NoCustomer struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Customer string `json:"customer"`
}

func (ch *CustomersHandler) GetCustomerHandler(c *fiber.Ctx) error {
	childCtx, span := tracer.StartSpan(c.Context(), "CustomersHandler.GetCustomerHandler")
	defer span.End()

	phoneNumber := c.Query("phoneNumber")
	merchantCode := c.Query("merchantCode")

	req := forms.GetRequest{
		MerchantCode: merchantCode,
		PhoneNumber:  phoneNumber,
	}
	value, err := ch.Service.GetCustomer(childCtx, req)
	if err != nil {
		ch.l.Errorf("Error when getting data : [%s]", err)
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.CommonError{
			Code:    fiber.StatusInternalServerError,
			Message: "internal server error",
			Error:   err.Error(),
		})
	}

	var customerDetails *models.CustomerDetailItem

	if value == nil {
		return c.Status(fiber.StatusOK).JSON(
			&NoCustomer{
				Code:     fiber.StatusOK,
				Message:  http.StatusText(fiber.StatusOK),
				Customer: "No Customer Found",
			})
	} else {
		customerDetails = &models.CustomerDetailItem{
			Id:           value.Code,
			MerchantCode: value.MerchantCode,
			PhoneNumber:  value.PhoneNumber,
			Email:        value.Email,
			Status:       value.Status,
		}
	}

	transactionId := childCtx.Value(constant.CtxTransactionId).(string)
	referenceNumber := childCtx.Value(constant.CtxReferenceNumber).(string)

	spanAttributes := map[string]string{
		constant.CtxTransactionId:   transactionId,
		constant.CtxReferenceNumber: referenceNumber,
	}

	tracer.SetSpanAttributes(span, spanAttributes)

	fmt.Println("Handler :", transactionId)
	fmt.Println("Handler :", referenceNumber)

	return c.Status(fiber.StatusOK).JSON(
		&models.CustomerDetailResponse{
			Code:     fiber.StatusOK,
			Message:  constant.Message_ok,
			Customer: *customerDetails,
		})
}
