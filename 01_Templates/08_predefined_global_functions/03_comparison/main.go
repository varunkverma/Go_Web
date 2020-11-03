package main

import (
	"log"
	"os"
	"text/template"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles("template1.gohtml"))
}

func main() {
	scores := struct {
		Score1 int
		Score2 int
	}{
		Score1: 7,
		Score2: 9,
	}
	err := templates.ExecuteTemplate(os.Stdout, "template1.gohtml", scores)
	if err != nil {
		log.Fatalln(err)
	}
}
