package config

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var templates map[string]*template.Template

//Compile view templates
func TemplateBuilder() {
	router := gin.Default()
	router.LoadHTMLGlob("Templates/*")
	router.GET("/user/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register",
		})
	})
	/*
		if templates == nil {
			templates = make(map[string]*template.Template)
		}
		templates["index"] = template.Must(template.ParseFiles("templates/index.html",
			"templates/base.html"))
		templates["about"] = template.Must(template.ParseFiles("templates/about.html",
			"templates/base.html"))
		templates["books"] = template.Must(template.ParseFiles("templates/books.html",
			"templates/base.html"))
		templates["contact"] = template.Must(template.ParseFiles("templates/contact.html",
			"templates/base.html"))
		templates["dash"] = template.Must(template.ParseFiles("templates/dashboard/dash.html",
			"templates/base.html"))
		templates["addBook"] = template.Must(template.ParseFiles("templates/dashboard/addBook.html",
			"templates/base.html"))
		templates["deleteBook"] = template.Must(template.ParseFiles("templates/dashboard/deleteBook.html",
			"templates/base.html"))
		templates["postBook"] = template.Must(template.ParseFiles("templates/dashboard/postBook.html",
			"templates/base.html"))
		templates["updateBook"] = template.Must(template.ParseFiles("templates/dashboard/updatBook.html",
			"templates/base.html"))*/

}

//Render templates for the given name, template definition and data object
func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
