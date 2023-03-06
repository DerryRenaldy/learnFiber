package models

type CustomerDetailResponse struct {
	Code     int                `json:"code"`
	Message  string             `json:"message"`
	Customer CustomerDetailItem `json:"customer"`
}

type CustomerDetailItem struct {
	Id           string `json:"id"`
	MerchantCode string `json:"merchantCode"`
	PhoneNumber  string `json:"phoneNumber"`
	Email        string `json:"email"`
	Status       int    `json:"status"`
}

type CreateCustomerResponse struct {
	Code     int                  `json:"code"`
	Message  string               `json:"message"`
	Customer CreateCustomerDetail `json:"customer"`
}

type CreateCustomerDetail struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
}

type CustomerItemUpdate struct {
	Code   string `json:"publicCustomerId"`
	Status int    `json:"status"`
}
