package main

import (
	"github.com/R0DR160HM/gin-api-rest/database"
	"github.com/R0DR160HM/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
