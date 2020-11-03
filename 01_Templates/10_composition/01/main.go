package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles("template1.gohtml"))
}

func main() {
	p1 := person{
		Name: "James Bond",
		Age:  40,
	}

	err := templates.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
