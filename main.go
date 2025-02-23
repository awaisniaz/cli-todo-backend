package main

import (
	"github.com/cli-todo/database"
	"github.com/cli-todo/routes"
)

func main() {
	// cmd.Execute()
	database.ConnectDatabase()
	routes := routes.SetupRoutes()
	routes.Run(":4000")
}
