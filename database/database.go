package database

import (
	"context"
	"fmt"
	"github.com/geraldywy/go-webserver-template/config"
	"github.com/geraldywy/go-webserver-template/models"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var DB *pg.DB

func InitDB(ctx context.Context) {
	fmt.Println("initializing db")
	db := pg.Connect(&pg.Options{
		Addr:                  config.Conf.DBHost,
		User:                  config.Conf.DBUser,
		Password:              config.Conf.DBPass,
		Database:              config.Conf.DBName,
	})
	if err := createSchema(db); err != nil {
		panic(fmt.Sprintf("failed to create schema, err: %v", err))
	}

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
	DB = db
}

func createSchema(db *pg.DB) error {
	schemas := []interface{}{
		(*models.User)(nil),
	}

	for _, schema := range schemas {
		err := db.Model(schema).CreateTable(&orm.CreateTableOptions{
			Varchar:       0,
			Temp:          false,
			IfNotExists:   true,
			FKConstraints: false,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func CloseDB() {
	DB.Close()
}