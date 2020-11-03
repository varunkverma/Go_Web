package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	templates, err := template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = templates.Execute(os.Stdout, nil)
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
