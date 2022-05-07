package main

import (
	"fmt"

	router "github.com/Miguel-Florian/E-School/Routers"
)

func main() {
	fmt.Println("This Site is building, it contains e-book for programming languages")
	router.InitAllRoutes() // contains all routes with middlewares
	// routers initialization
}
