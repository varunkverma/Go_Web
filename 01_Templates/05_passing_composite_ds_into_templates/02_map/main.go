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
	foodPrices := map[string]int{
		"Pizza":        250,
		"Dragon Rolls": 100,
		"Momos":        60,
	}

	err := templates.ExecuteTemplate(os.Stdout, "template.gohtml", foodPrices)
	if err != nil {
		log.Fatalln(err)
	}
	err = templates.ExecuteTemplate(os.Stdout, "template2.gohtml", foodPrices)
	if err != nil {
		log.Fatalln(err)
	}
}
