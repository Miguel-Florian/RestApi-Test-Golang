package main

import (
	"fmt"

	router "github.com/Miguel-Florian/E-School/Routers"
)

func main() {
	fmt.Println("Starting server ...")
	//controllers.InitialBookMigration()
	router.InitAllRoutes() // Initialize all routes
	//config.TemplateBuilder()
}
