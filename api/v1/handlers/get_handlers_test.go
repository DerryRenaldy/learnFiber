package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/DerryRenaldy/learnFiber/api/v1/services"
	"github.com/DerryRenaldy/learnFiber/constant"
	"github.com/DerryRenaldy/learnFiber/entity"
	"github.com/DerryRenaldy/learnFiber/server/middleware"
	"github.com/DerryRenaldy/logger/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCustomersHandler_GetCustomerHandler(t *testing.T) {
	l := logger.New("test", "test", "debug")
	ctrl := gomock.NewController(t)
	serviceMock := services.NewMockIService(ctrl)

	ch := &CustomersHandler{
		l:       l,
		Service: serviceMock,
	}

	actualPath := "/api/v1"

	app := fiber.New(fiber.Config{
		Immutable: true,
	})
	app.Get(actualPath, middleware.ValidateHeaderMiddleware(), ch.GetCustomerHandler)

	tests := []struct {
		name                  string
		method                string
		path                  string
		statusCode            int
		body                  string
		requestBody           map[string]interface{}
		handlerMethodName     string
		handlerToBeCalledWith []interface{}
		requestHeaders        map[string]string
		setMock               func()
	}{
		{
			name:              "GET endpoint success",
			method:            http.MethodGet,
			path:              "/api/v1",
			statusCode:        200,
			body:              `{"code":200,"message":"ok","customer":{"id":"TEST-1cfds3","merchantCode":"TEST1","phoneNumber":"0852345236483763","email":"test@test.com","status":1}}`,
			requestBody:       nil,
			handlerMethodName: "GetCustomerHandler",
			setMock: func() {
				serviceMock.EXPECT().GetCustomer(gomock.Any(), gomock.Any()).Return(
					&entity.Customer{
						ID:               1,
						Code:             "TEST-1cfds3",
						PublicCustomerId: "TEST-1cfds3",
						MerchantCode:     "TEST1",
						PhoneNumber:      "0852345236483763",
						Email:            "test@test.com",
						Status:           1,
					}, nil,
				)
			},
		},
		{
			name:              "GET endpoint failed",
			method:            http.MethodGet,
			path:              "/api/v1",
			statusCode:        200,
			body:              `{"code":200,"message":"OK","customer":"No Customer Found"}`,
			requestBody:       nil,
			handlerMethodName: "GetCustomerHandler",
			setMock: func() {
				serviceMock.EXPECT().GetCustomer(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
		},
		{
			name:              "GET missing mandatory fields",
			method:            http.MethodGet,
			path:              "/api/v1",
			statusCode:        500,
			body:              `{"code":500,"message":"internal server error","error":"Internal Server Error"}`,
			requestBody:       nil,
			handlerMethodName: "GetCustomerHandler",
			setMock: func() {
				serviceMock.EXPECT().GetCustomer(gomock.Any(), gomock.Any()).Return(nil, fiber.ErrInternalServerError)
			},
		},
		{
			name:              "GET wrong endpoint",
			method:            http.MethodGet,
			path:              "/api/v2",
			statusCode:        404,
			body:              `Cannot GET /api/v2`,
			requestBody:       nil,
			handlerMethodName: "GetCustomerHandler",
			setMock:           func() {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setMock()
			// Create and send request
			rbody, _ := json.Marshal(tt.requestBody)
			request := httptest.NewRequest(tt.method, tt.path, bytes.NewReader(rbody))
			request.Header.Add(`Content-Type`, `application/json`)
			ctx := request.Context()
			ctx = context.WithValue(ctx, constant.CtxTransactionId, "transactionIdTest")
			ctx = context.WithValue(ctx, constant.CtxReferenceNumber, "referenceNumberTest")

			response, _ := app.Test(request)

			// Validating
			// Status Code
			statusCode := response.StatusCode
			if diff := cmp.Diff(statusCode, tt.statusCode); diff != "" {
				t.Fatalf("\t%s\tStatusCode was incorrect, got: %d, want: %d.", constant.Failed, tt.statusCode,
					statusCode)
			}
			t.Logf("\t%s\tShould get statusCode is %v", constant.Succeed, statusCode)

			// Response Body
			body, _ := io.ReadAll(response.Body)
			actual := string(body)
			if diff := cmp.Diff(actual, tt.body); diff != "" {
				t.Fatalf("\t%s\tBody was incorrect, got: %v, want: %v", constant.Failed, tt.body, actual)
			}
			t.Logf("\t%s\tShould get body is %v", constant.Succeed, actual)
		})
	}
}
