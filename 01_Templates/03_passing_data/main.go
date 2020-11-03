package main

import (
	"log"
	"os"
	"text/template"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	err := templates.ExecuteTemplate(os.Stdout, "template.gohtml", "Go Web Dev")
	if err != nil {
		log.Fatalln(err)
	}
}
