package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"goshop-api/app/route"
	"goshop-api/util"
	"os"
)

func main() {

	//load env
	if err := godotenv.Load(); err != nil {
		fmt.Println("failed load env file")
		panic(err)
	}

	//initialize sql database
	util.InitializeSqlDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		)

	//serve the route link and start gin
	route.ServeRoutes()
}
