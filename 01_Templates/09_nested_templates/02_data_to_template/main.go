package main

import (
	"log"
	"os"
	"text/template"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	err := templates.ExecuteTemplate(os.Stdout, "index.gohtml", "Hakuna Matata")
	if err != nil {
		log.Fatalln(err)
	}
}
