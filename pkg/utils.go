package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/mail"
	"regexp"
	"strings"
)

func ValidateTxnIdRefNum(ctx context.Context, transactionId, referenceNumber string) error {
	var regex = regexp.MustCompile(`^[a-zA-Z0-9]{32}$`)
	if !regex.MatchString(transactionId) {
		return fiber.ErrBadRequest
	}

	if !regex.MatchString(referenceNumber) {
		return fiber.ErrBadRequest
	}
	return nil
}

func ValidateMerchantCode(ctx context.Context, merchantCode string) error {
	var regex = regexp.MustCompile(`^[A-Z0-9]{6}$`)
	if !regex.MatchString(merchantCode) {
		return fiber.ErrBadRequest
	}
	return nil
}

func ValidatePhone(PhoneNumber string) (string, error) {
	res := PhoneNumber
	if !strings.HasPrefix(PhoneNumber, "62") {
		PhoneNumber = strings.TrimPrefix(PhoneNumber, "0")
		PhoneNumber = fmt.Sprintf("62%s", PhoneNumber)
	}
	number := strings.TrimPrefix(PhoneNumber, "62")
	if len(number) < 9 || len(number) > 13 {
		return res, errors.New("Invalid Phone Number")
	}
	return PhoneNumber, nil
}

func ValidateEmail(email string) bool {
	_, emailErr := mail.ParseAddress(email)
	return emailErr == nil
}
