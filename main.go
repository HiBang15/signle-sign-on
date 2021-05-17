package main

import (
	"fmt"
	"github.com/HiBang15/signle-sign-on/api/rest/router"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	// load config from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	//run Rest API
	fmt.Println("Start run REST API .....")
	router.Start(os.Getenv("ENVIRONMENT"))
}