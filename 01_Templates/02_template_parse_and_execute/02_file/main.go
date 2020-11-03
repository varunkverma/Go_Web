package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	template, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln("err")
	}

	file, err := os.Create("index.html")
	defer file.Close()

	err = template.Execute(file, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
