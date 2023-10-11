package main

import (
	"github.com/mcsans/go-jwt/database"
	"github.com/mcsans/go-jwt/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}