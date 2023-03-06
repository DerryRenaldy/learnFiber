package handlers

import (
	"fmt"
	"github.com/DerryRenaldy/learnFiber/errors"
	"github.com/DerryRenaldy/learnFiber/forms"
	"github.com/DerryRenaldy/learnFiber/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (ch *CustomersHandler) UpdateCustomerHandler(c *fiber.Ctx) error {
	functionName := "CustomersHandler.CreateCustomerHandler"
	ctx := c.Context()

	fmt.Println(c.Method())

	if c.Method() != http.MethodPatch {
		err := fmt.Errorf("[%s : Used Wrong Request Method]", functionName)
		ch.l.Error(err.Error())
		return err
	}

	payload := new(forms.UpdateRequest)
	if err := c.BodyParser(payload); err != nil {
		ch.l.Errorf("[%s - c.BodyParser] : %s", functionName, err)
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.CommonError{
			Code:    fiber.StatusInternalServerError,
			Message: fiber.ErrInternalServerError.Error(),
			Error:   err.Error(),
		})
	}

	customerObj, err := ch.Service.UpdateByCustomerId(ctx,
		forms.UpdateRequest{
			PublicCustomerId: payload.PublicCustomerId,
			Status:           payload.Status,
		},
	)
	if err != nil {
		ch.l.Errorf("[%s - ch.Service.UpdateByCustomerId] : %s", functionName, err)
		return c.Status(fiber.StatusInternalServerError).JSON(&errors.CommonError{
			Code:    fiber.StatusInternalServerError,
			Message: fiber.ErrInternalServerError.Error(),
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&models.CustomerItemUpdate{
		Code:   customerObj.Code,
		Status: customerObj.Status,
	})
}
