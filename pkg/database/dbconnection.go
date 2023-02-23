package database

import (
	"database/sql"
	"github.com/DerryRenaldy/logger/logger"
	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	l logger.ILogger
}

type DBConnection interface {
	DBConnect()
}

func NewDatabaseConnection(logger logger.ILogger) *Connection {
	return &Connection{l: logger}
}

func (db *Connection) DBConnect() *sql.DB {
	dbConn, errConn := sql.Open("mysql",
		"root:root@unix(/cloudsql/kubernetes-gcp-378606:asia-southeast2:mysql-diber-project)/customers")

	if errConn != nil {

		db.l.Fatalf("Error while connecting database. err= %v", errConn.Error())
		return nil
	}
	errPing := dbConn.Ping()
	if errPing != nil {
		db.l.Fatalf("Error while connecting database. err= %v", errPing.Error())
		return nil
	}
	return dbConn
}
