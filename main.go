package main

import (
	"database/sql"
	"fmt"

	"github.com/gabelev/get_heard/lib/data"
	"github.com/gabelev/get_heard/lib/routes"
	"github.com/gabelev/get_heard/lib/settings"
)

func init() {
	settings.SetConfig()
}

func main() {
	dbHandler, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=disable",
			settings.Config.DbUser,
			settings.Config.DbPassword,
			settings.Config.DbName,
		),
	)
	if err != nil {
		panic(err)
	}
	db, err := data.New(dbHandler)
	if err != nil {
		panic(err)
	}
	router := routes.NewServer()
	service := routes.New(
		"localhost",
		settings.Config.AppPort,
		db,
		router,
	)
	err = service.Start()
	if err != nil {
		panic(err)
	}

}
