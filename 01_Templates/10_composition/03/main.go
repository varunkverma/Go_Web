package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles("template1.gohtml"))
}

func main() {

	year := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{
					Number: "CS-101",
					Name:   "Intro to OS",
					Units:  "3",
				},
				course{
					Number: "CS-102",
					Name:   "Intro to Network",
					Units:  "4",
				},
				course{
					Number: "CS-103",
					Name:   "Intro to DBMS",
					Units:  "2",
				},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{
					Number: "CS-201",
					Name:   "Intro to CPP",
					Units:  "1",
				},
				course{
					Number: "CS-202",
					Name:   "Intro to GO",
					Units:  "4",
				},
				course{
					Number: "CS-203",
					Name:   "Intro to EC",
					Units:  "2",
				},
			},
		},
		Summer: semester{
			Term: "Summer",
			Courses: []course{
				course{
					Number: "CS-301",
					Name:   "Intro to DSA",
					Units:  "5",
				},
				course{
					Number: "CS-302",
					Name:   "Intro to OOPS",
					Units:  "2",
				},
				course{
					Number: "CS-303",
					Name:   "Intro to EE",
					Units:  "2",
				},
			},
		},
	}

	err := templates.Execute(os.Stdout, year)
	if err != nil {
		log.Fatalln(err)
	}
}
