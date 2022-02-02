package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DbManager interface {
	NewConnection() (*sql.DB, error)
}

type DbManagerImpl struct {

}

func (d DbManagerImpl) NewConnection() (*sql.DB, error) {

	//TODO: Get config from viper
	address := "localhost"
	port := "3306"
	database := "db"
	user := "user"
	password := "12345!"
	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		user,
		password,
		address,
		port,
		database,
		param)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("error conecting to database %v", err)
	}

	return db, nil
}

