package main

import (
	"github.com/FMichetti/api-go-gin/database"
	_ "github.com/FMichetti/api-go-gin/docs"
	"github.com/FMichetti/api-go-gin/routes"
)

// @title Test API
// @version 1.0
// @description A Test API using GO, Gin and Gorm

// @host localhost:8080

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
