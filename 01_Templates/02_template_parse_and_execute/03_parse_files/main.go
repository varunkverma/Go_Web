package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	templates, err := template.ParseFiles("template1.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = templates.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	templates, err = templates.ParseFiles("template2.gohtml", "template3.gohtml")
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

	err = templates.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
