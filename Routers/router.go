package routers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitAllRoutes() {
	router := mux.NewRouter().StrictSlash(false)
	router = SetUserRoutes(router) // init routes for user
	router = SetBookRoutes(router) // init routes for book
	
	log.Fatal(http.ListenAndServe(":8000",router))
}
