package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lgarcia93/shoplist/internal/config"
	"log"
)

type DbManager interface {
	NewConnection(env string) (*sql.DB, error)
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

	c.ExecContext(ctx, dbDDL)

	return nil
}

func (d DbManagerImpl) NewConnection(env string) (*sql.DB, error) {
	param := "parseTime=true"

	var secretData config.SecretData

	if env == "prod" {
		secretData = config.GetSecret()
	} else {
		secretData = config.SecretData{
			MySQLUser:               "user",
			MySQLPass:               "12345!",
			MySQLPort:               3306,
			MySQLAddress:            "localhost",
			MySQLDatabaseName:       "db",
			MySQLInstanceIdentifier: "",
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
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
		log.Fatal("error executing DDL %v", err)
	}

	return db, nil
}
