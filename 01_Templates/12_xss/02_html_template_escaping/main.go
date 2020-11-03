package main

import (
	"html/template"
	"log"
	"os"
)

type page struct {
	Title   string
	Heading string
	Input   string
}

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles("template1.gohtml"))
}

func main() {
	home := page{
		Title:   "Danger Escaped",
		Heading: "Html escaped with html/template",
		Input:   `<script>alert("Boo!")</script>`,
	}

	err := templates.Execute(os.Stdout, home)
	if err != nil {
		log.Fatalln(err)
	}
}
