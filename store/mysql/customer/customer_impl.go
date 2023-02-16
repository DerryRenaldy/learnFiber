package customer

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/DerryRenaldy/learnFiber/entity"
	"github.com/DerryRenaldy/logger/logger"
	"github.com/go-sql-driver/mysql"
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
	functionName := "customerRepoImpl.GetByPhoneNumber"
	res := new(entity.Customer)
	err := o.DB.QueryRow(
		"select c.id,c.code,c.status,ce.email,cm.merchantCode,cp.phoneNumber from customers.customers c inner Join customers.customers_merchants cm on cm.customerCode=c.code join customers.customers_phone_numbers cp on c.code=cp.customerCode join customers.customers_emails ce on c.code=ce.customercode where cp.phoneNumber=? and cm.merchantCode=?;",
		obj.PhoneNumber, obj.MerchantCode,
	).Scan(&res.ID, &res.Code, &res.Status, &res.Email, &res.MerchantCode, &res.PhoneNumber)
	if err != nil {
		o.l.Errorf("[%s - Eo.DB.QueryRow] : %s", functionName, err)
		return res, err
	}

	resByte, _ := json.Marshal(res)
	o.l.Debugf("Query Result : %s", resByte)
	return res, nil
}

func (o *customerRepoImpl) SaveCustomer(ctx context.Context, req *entity.Customer) (int64, error) {
	functionName := "customerRepoImpl.SaveCustomer"
	var customerId int64

	stmt, err := o.DB.Prepare("INSERT into customers.customers SET code=?, status=?;")
	if err != nil {
		o.l.Errorf("[%s - o.DB.Prepare] : %s", functionName, err)
		return customerId, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(req.Code, req.Status)
	if err != nil {
		o.l.Errorf("[%s - stmt.Exec]: %s", functionName, err)
		var mysqlErr = new(mysql.MySQLError)
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			o.l.Errorf("[%s - duplicate error ]: %s", functionName, mysqlErr)
			return customerId, errors.New("Duplicate_Entry")
		}
		return customerId, err
	}
	customerId, _ = result.LastInsertId()

	return customerId, nil
}

func (o *customerRepoImpl) SaveMerchant(ctx context.Context, req *entity.Customer) (bool, error) {
	functionName := "customerRepoImpl.SaveMerchant"

	stmt, err := o.DB.Prepare("INSERT into customers.customers_merchants SET customerCode=?, merchantCode=?;")
	if err != nil {
		o.l.Errorf("[%s - o.DB.Prepare]: %s", functionName, err)
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Code, req.MerchantCode)
	if err != nil {
		o.l.Errorf("[%s - stmt.Exec]: %s", functionName, err)
		return false, err
	}

	return true, nil
}

func (o *customerRepoImpl) SavePhoneNumber(ctx context.Context, req *entity.Customer) (bool, error) {
	functionName := "customerRepoImpl.SavePhoneNumber"

	stmt, err := o.DB.Prepare("INSERT into customers.customers_phone_numbers SET customerCode=?, phoneNumber=?;")
	if err != nil {
		o.l.Errorf("[%s - o.DB.Prepare]: %s", functionName, err)
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Code, req.PhoneNumber)
	if err != nil {
		o.l.Errorf("[%s - stmt.Exec]: %s", functionName, err)
		return false, err
	}

	return true, nil
}

func (o *customerRepoImpl) SaveEmail(ctx context.Context, req *entity.Customer) (bool, error) {
	functionName := "customerRepoImpl.SaveEmail"

	stmt, err := o.DB.Prepare("INSERT into customers.customers_emails SET customerCode=?, email=?;")
	if err != nil {
		o.l.Errorf("[%s - o.DB.Prepare]: %s", functionName, err)
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Code, req.Email)
	if err != nil {
		o.l.Errorf("[%s - stmt.Exec]: %s", functionName, err)
		return false, err
	}

	return true, nil
}
