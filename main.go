package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"goshop-api/app/route"
	_ "goshop-api/util" //initialize db
)

func main() {

	//load env
	if err := godotenv.Load(); err != nil {
		fmt.Println("failed load env file")
		panic(err)
	}

	//serve the route link and start gin
	route.ServeRoutes()
}
