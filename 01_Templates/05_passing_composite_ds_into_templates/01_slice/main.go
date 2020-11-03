package main

import (
	"log"
	"os"
	"text/template"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles("template.gohtml", "template2.gohtml"))
}

func main() {
	favouriteFood := []string{
		"Pizza",
		"Dragon Rolls",
		"Momos",
	}

	err := templates.ExecuteTemplate(os.Stdout, "template.gohtml", favouriteFood)
	if err != nil {
		log.Fatalln(err)
	}
	err = templates.ExecuteTemplate(os.Stdout, "template2.gohtml", favouriteFood)
	if err != nil {
		log.Fatalln(err)
	}
}
