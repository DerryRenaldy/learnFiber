package constant

const (
	HeaderContentType   = "Content-Type"
	HeaderAccept        = "Accept"
	MIMEApplicationJson = "application/json"

	ResponseTimeLayout = "20060102150405"

	CtxTransactionId   = "transactionId"
	CtxReferenceNumber = "referenceNumber"

	Customer   = "customers"
	Message_ok = "ok"

	ErrTransactionId    = "PRECONDITION_FAILED_TRANSACTION_ID"
	ErrReferenceNumber  = "PRECONDITION_FAILED_REFERENCE_NUMBER"
	ErrMerchantCode     = "PRECONDITION_FAILED_MERCHANT_CODE"
	ErrNotFoundCustomer = "NOT_FOUND_CUSTOMER"
	ErrSqlError         = "SQL_ERROR"
)
