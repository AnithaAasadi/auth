package main

import (
	"task1/routes"
	"task1/database"


)

func main() {

	
	database.Setup()
	routes.Routes()
	
}
