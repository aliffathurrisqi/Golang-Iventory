package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBconnection() (*sql.DB, error){
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "db_iventory"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}