package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/spani73/go-ecommerce-api/cmd/api"
	"github.com/spani73/go-ecommerce-api/config"
	"github.com/spani73/go-ecommerce-api/db"
)

func main() {
	db , err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)
	server := api.NewAPIServer(":8080" , db)
	if err := server.Run(); err  != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB){
	err := db.Ping()
	if err != nil{
		log.Fatal(err)
	}

	log.Println("Database is successfully connected")
}