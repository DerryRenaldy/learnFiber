package customer

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/DerryRenaldy/learnFiber/entity"
	"github.com/DerryRenaldy/logger/logger"
)

type customerRepoImpl struct {
	l  logger.ILogger
	DB *sql.DB
}

func NewCustomerStoreImpl(logger logger.ILogger, db *sql.DB) *customerRepoImpl {
	return &customerRepoImpl{
		l:  logger,
		DB: db,
	}
}

func (o *customerRepoImpl) GetByPhoneNumber(ctx context.Context, obj *entity.Customer) (*entity.Customer, error) {
	res := new(entity.Customer)
	err := o.DB.QueryRow(
		"select c.id,c.code,c.status,cm.merchantCode,cp.phoneNumber from customers c inner Join customers_merchants cm on cm.customerCode=c.code join customers_phone_numbers cp on c.code=cp.customerCode where cp.phoneNumber=? and cm.merchantCode=?",
		obj.PhoneNumber, obj.MerchantCode,
	).Scan(&res.ID, &res.Code, &res.Status, &res.MerchantCode, &res.PhoneNumber)
	if err != nil {
		o.l.Errorf("Error getting data from database : %s", err)
		return res, err
	}

	resByte, _ := json.Marshal(res)
	o.l.Debugf("Query Result : %s", resByte)
	return res, nil
}
