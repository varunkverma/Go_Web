package main

import (
	"log"
	"os"
	"text/template"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	err := templates.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = templates.ExecuteTemplate(os.Stdout, "template3.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = templates.ExecuteTemplate(os.Stdout, "template1.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = templates.ExecuteTemplate(os.Stdout, "template2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
