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

type user struct {
	Name  string
	Motto string
	Admin bool
}

func main() {
	users := []user{
		user{
			Name:  "Buddha",
			Motto: "The belief of no beliefs",
			Admin: false,
		},
		user{
			Name:  "Gandhi",
			Motto: "Be the change",
			Admin: true,
		},
		user{
			Name:  "",
			Motto: "Nobody",
			Admin: true,
		},
	}
	err := templates.ExecuteTemplate(os.Stdout, "template1.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
