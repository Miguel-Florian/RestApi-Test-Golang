package routers

import (
	controllers "github.com/Miguel-Florian/E-School/Controllers"
	"github.com/gorilla/mux"
)

func SetBookRoutes(router *mux.Router) *mux.Router {
	bookRoutes := mux.NewRouter()
	bookRoutes.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	bookRoutes.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	bookRoutes.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	bookRoutes.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	bookRoutes.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
	//bookRoutes.HandleFunc("/notes/tasks/{id}", controllers.GetNotesByTask).Methods("GET")
	/*router.PathPrefix("/notes").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(noteRouter),
	))*/
	return router
}
