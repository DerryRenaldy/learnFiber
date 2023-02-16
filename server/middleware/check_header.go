package middleware

import (
	"fmt"
	"github.com/DerryRenaldy/learnFiber/constant"
	"github.com/gofiber/fiber/v2"
)

func ValidateHeaderMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		c.Set(constant.HeaderContentType, constant.MIMEApplicationJson)

		transactionId := c.Query("transactionId")
		referenceNumber := c.Query("referenceNumber")

		ctx.SetUserValue(constant.CtxTransactionId, transactionId)
		ctx.SetUserValue(constant.CtxReferenceNumber, referenceNumber)

		method := c.Method()
		if method != "GET" {
			type RequestBody struct {
				TransactionId   string `json:"transactionsId"`
				ReferenceNumber string `json:"referenceNumber"`
			}

			requestBody := RequestBody{}

			err := c.BodyParser(&requestBody)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON("error in middleware")
			}
		}

		fmt.Println("middleware :", transactionId)
		fmt.Println("middleware :", referenceNumber)

		return c.Next()
	}
}
