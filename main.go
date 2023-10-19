package main

import (
	"github.com/FMichetti/api-go-gin/config"
	_ "github.com/FMichetti/api-go-gin/docs"
	"github.com/FMichetti/api-go-gin/routes"
)

// @title Test API
// @version 1.0
// @description Api para criação, correção e execução de Avaliações Escolares

// @host localhost:8080

func main() {
	config.ConectaComBancoDeDados()
	routes.Run()
}
