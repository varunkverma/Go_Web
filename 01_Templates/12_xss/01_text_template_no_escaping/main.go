package main

import (
	"log"
	"os"
	"text/template"
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
		Title:   "Nothing Escaped",
		Heading: "Nothing escaped with text/template",
		Input:   `<script>alert("Boo!")</script>`,
	}

	err := templates.Execute(os.Stdout, home)
	if err != nil {
		log.Fatalln(err)
	}
}
