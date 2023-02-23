package database

import (
	"database/sql"
	"fmt"
	"github.com/DerryRenaldy/logger/logger"
	_ "github.com/go-sql-driver/mysql"
	"os"
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
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASS")
	connectionType := os.Getenv("CONTYPE")
	connectionName := os.Getenv("CONNAME")
	database := os.Getenv("DATABASE")

	dbConn, errConn := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s)/%s", username, password, connectionType, connectionName, database))

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
