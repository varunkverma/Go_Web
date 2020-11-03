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
	data := struct {
		Words []string
		LName string
	}{
		Words: xs,
		LName: "Bond",
	}

	err := templates.ExecuteTemplate(os.Stdout, "template1.gohtml", data)
	if err != nil {
		log.Fatalln(xs)
	}
}
