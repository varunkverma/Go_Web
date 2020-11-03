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

type doubleZero struct {
	person
	LicenseToKill bool
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
	doubleZero := doubleZero{
		person:        p1,
		LicenseToKill: false,
	}

	err := templates.Execute(os.Stdout, doubleZero)
	if err != nil {
		log.Fatalln(err)
	}
}
