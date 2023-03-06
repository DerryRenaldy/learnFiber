package forms

type UpdateRequest struct {
	TransactionId    string `json:"transactionId"`
	ReferenceNumber  string `json:"referenceNumber"`
	PublicCustomerId string `json:"publicCustomerId"`
	Status           int    `json:"status"`
}
