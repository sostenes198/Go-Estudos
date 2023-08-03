package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome string
	Email string
}

func main() {
	fmt.Println("HTML")

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"João", "joao@com.br"}
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Inicializando servidor")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
