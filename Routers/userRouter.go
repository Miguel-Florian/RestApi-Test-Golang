package routers

import (
	controllers "github.com/Miguel-Florian/E-School/Controllers"
	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	userRoutes := mux.NewRouter()
	userRoutes.HandleFunc("/user/register", controllers.CreateUser).Methods("POST")
	userRoutes.HandleFunc("/user/login", controllers.LoginUser).Methods("PUT")
	userRoutes.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	userRoutes.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	userRoutes.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	//bookRoutes.HandleFunc("/notes/tasks/{id}", controllers.GetNotesByTask).Methods("GET")
	/*router.PathPrefix("/notes").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(noteRouter),
	))*/
	return router
}
