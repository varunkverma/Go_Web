package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type villian struct {
	Name  string
	Motto string
}

var templates *template.Template

var functionMap = template.FuncMap{
	"upCase":     strings.ToUpper,
	"firstThree": firstThree,
}

func init() {
	templates = template.Must(template.New("").Funcs(functionMap).ParseFiles("template1.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	villians := []villian{
		villian{
			Name:  "Darth Vadar",
			Motto: "Come, join the dark side",
		},
		villian{
			Name:  "Mega Mind",
			Motto: "Ollo",
		},
		villian{
			Name:  "Darth Lord",
			Motto: "Power, infite power",
		},
	}
	err := templates.ExecuteTemplate(os.Stdout, "template1.gohtml", villians)
	if err != nil {
		log.Fatalln(err)
	}
}
