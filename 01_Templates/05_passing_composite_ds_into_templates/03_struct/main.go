package main

import (
	"log"
	"os"
	"text/template"
)

type villian struct {
	Name  string
	Motto string
}

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles("template1.gohtml", "template2.gohtml"))
}

func main() {
	darthVadar := villian{
		Name:  "Darth Vadar",
		Motto: "Come, Join the Dark Side",
	}
	err := templates.ExecuteTemplate(os.Stdout, "template1.gohtml", darthVadar)
	if err != nil {
		log.Fatalln(err)
	}
	err = templates.ExecuteTemplate(os.Stdout, "template2.gohtml", darthVadar)
	if err != nil {
		log.Fatalln(err)
	}
}
