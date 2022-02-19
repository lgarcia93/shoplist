package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lgarcia93/shoplist/internal/config"
	"io/ioutil"
	"log"
)

type DbManager interface {
	NewConnection() (*sql.DB, error)
	ExecDDL(db *sql.DB) error
}

type DbManagerImpl struct {
}

func (d DbManagerImpl) ExecDDL(db *sql.DB) error {
	ctx := context.Background()

	c, err := db.Conn(ctx)

	if err != nil {
		return err
	}

	buffer, err := ioutil.ReadFile("ddl.sql")

	if err != nil {
		return err
	}

	ddlContent := string(buffer)

	c.ExecContext(ctx, ddlContent)

	return nil
}

func (d DbManagerImpl) NewConnection() (*sql.DB, error) {
	//address := "localhost"
	//port := "3306"
	//database := "db"
	//user := "user"
	//password := "12345!"
	param := "parseTime=true"

	secretData, err := config.GetSecret()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		secretData.MySQLUser,
		secretData.MySQLPass,
		secretData.MySQLAddress,
		secretData.MySQLPort,
		secretData.MySQLDatabaseName,
		param)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("error conecting to database %v", err)
	}

	db.Conn(context.Background())

	err = d.ExecDDL(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}
