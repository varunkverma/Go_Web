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
	xs := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
	}
	err := templates.ExecuteTemplate(os.Stdout, "template1.gohtml", xs)
	if err != nil {
		log.Fatalln(xs)
	}
}
