package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yosefalemu/GoCrudChallange/initializers"
	"github.com/yosefalemu/GoCrudChallange/server"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	// migrations.Migrate(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
}

func main() {
	if initializers.DB == nil {
		log.Fatal("Database connection is not initialized")
	}
	s := server.NewServer()
	s.Initialize()
	serverPort := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	s.Run(serverPort)
}
