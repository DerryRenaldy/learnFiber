package entity

import "github.com/DerryRenaldy/learnFiber/constant"

type Customer struct {
	ID               int64  `json:"id"`
	Code             string `json:"code"`
	PublicCustomerId string `json:"publicCustomerId"`
	MerchantCode     string `json:"merchantCode"`
	PhoneNumber      string `json:"phoneNumber"`
	Email            string `json:"email,omitempty"`
	Status           int    `json:"status"`
}

func (c Customer) Tablename() string {
	return constant.Customer
}
