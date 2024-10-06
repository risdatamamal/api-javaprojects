package main

import (
	"fmt"

	"github.com/risdatamamal/api-javaprojects/config"
	"github.com/risdatamamal/api-javaprojects/database"
	"github.com/risdatamamal/api-javaprojects/router/v1"
)

func main() {
	r := router.StartApp()
	err := database.StartDB()

	if err != nil {
		fmt.Println("Error starting database: ", err)
		return
	}

	r.Run(config.SERVER_PORT)
}
