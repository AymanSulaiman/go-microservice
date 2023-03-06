package database

import (
	"database/sql"
	"fmt"
	"go-microservice/config"
	"log"
)

var DB *sql.DB

func Conn() {
	conf := config.Load()

	username := conf.Database.Username
	password := conf.Database.Password
	host := conf.Database.Host
	port := conf.Database.Port
	DBName := conf.Database.DBName

	var err error

	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", username, password, host, port, DBName))

	if err != nil {
		panic(err)
	}

	// pinging the server
	err = DB.Ping()

	if err != nil {
		log.Fatal(err)
	}
}
