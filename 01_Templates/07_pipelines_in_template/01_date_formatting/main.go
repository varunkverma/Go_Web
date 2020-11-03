package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var templates *template.Template

func monthDayYear(t time.Time) string {
	return t.Format("02-01-06")
}

var functionMap = template.FuncMap{
	"mdy": monthDayYear,
}

func init() {
	templates = template.Must(template.New("").Funcs(functionMap).ParseFiles("template1.gohtml"))
}

func main() {
	err := templates.ExecuteTemplate(os.Stdout, "template1.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
