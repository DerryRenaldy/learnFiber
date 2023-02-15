package handlers

import (
	"fmt"
	"github.com/DerryRenaldy/learnFiber/constant"
	"github.com/DerryRenaldy/learnFiber/forms"
	"github.com/DerryRenaldy/learnFiber/models"
	"github.com/gofiber/fiber/v2"
)

type NoCustomer struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (ch *CustomersHandler) GetCustomerHandler(c *fiber.Ctx) error {
	phoneNumber := c.Query("phoneNumber")
	merchantCode := c.Query("merchantCode")

	req := forms.GetRequest{
		MerchantCode: merchantCode,
		PhoneNumber:  phoneNumber,
	}
	value, err := ch.Service.GetCustomer(c.Context(), req)
	if err != nil {
		ch.l.Errorf("Error when getting data : [%s]", err)
		c.Status(fiber.StatusInternalServerError).JSON(fmt.Sprintf(`{"code": %d,"message": %s}`, fiber.StatusInternalServerError, err.Error()))
		return err
	}

	var customerDetails *models.CustomerDetailItem

	if value == nil {
		return c.Status(fiber.StatusCreated).JSON(
			&NoCustomer{
				Code:    fiber.StatusOK,
				Message: "No Customer Found",
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

	return c.Status(fiber.StatusCreated).JSON(
		&models.CustomerDetailResponse{
			Code:     fiber.StatusOK,
			Message:  constant.Message_ok,
			Customer: *customerDetails,
		})

}
