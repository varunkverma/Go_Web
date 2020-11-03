package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var templates *template.Template

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var functionMap = template.FuncMap{
	"fdbl":  double,
	"fsqr":  square,
	"fsqrt": sqRoot,
}

func init() {
	templates = template.Must(template.New("").Funcs(functionMap).ParseFiles("template1.gohtml"))
}

func main() {
	err := templates.ExecuteTemplate(os.Stdout, "template1.gohtml", 300)
	if err != nil {
		log.Fatalln(err)
	}
}
