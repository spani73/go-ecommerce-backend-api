package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spani73/go-ecommerce-api/config"
	"github.com/spani73/go-ecommerce-api/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	driver, _ := sql

	m, err := migrate.NewWithDatabaseInstance(
		"C:\\Personal\\Code\\Project\\Go\\go-ecommerce-backend-api\\cmd\\migrate\\migrations",
		"mysql",
		driver,
	)

}
