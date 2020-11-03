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

func (p person) SomeProcessing() int {
	return 7
}
func (p person) AgeDbl() int {
	return p.Age * 2
}
func (p person) TakeArgs(x int) int {
	return x * 2
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
