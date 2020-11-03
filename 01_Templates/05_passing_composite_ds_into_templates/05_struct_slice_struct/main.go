package main

import (
	"log"
	"os"
	"text/template"
)

var templates *template.Template

type villian struct {
	Name  string
	Motto string
}
type hero struct {
	Name  string
	Power string
}

func init() {
	templates = template.Must(template.ParseFiles("template1.gohtml", "template2.gohtml"))
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
	heroes := []hero{
		hero{
			Name:  "All Might",
			Power: "Super Strength",
		},
		hero{
			Name:  "Flash",
			Power: "Speed Force",
		},
		hero{
			Name:  "Luffy",
			Power: "Rubber man",
		},
	}

	characters := struct {
		Heroes   []hero
		Villians []villian
	}{
		Heroes:   heroes,
		Villians: villians,
	}

	err := templates.ExecuteTemplate(os.Stdout, "template1.gohtml", characters)
	if err != nil {
		log.Fatalln(err)
	}

	err = templates.ExecuteTemplate(os.Stdout, "template2.gohtml", characters)
	if err != nil {
		log.Fatalln(err)
	}

}
