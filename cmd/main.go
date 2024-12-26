package main

import (
	"deto_api/internals/database"
	"deto_api/internals/routes"
)

func main() {

	database.GetDB()
	database.Automigrate()

	r := routes.SetupRutes()
	
	r.Run()
}
