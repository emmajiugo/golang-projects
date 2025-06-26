package main

import (
	"log"

	"github.com/emmajiugo/ecom/cmd/api"
	"github.com/emmajiugo/ecom/config"
	"github.com/emmajiugo/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main()  {
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
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	// Initialize the API server with the database connection
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}