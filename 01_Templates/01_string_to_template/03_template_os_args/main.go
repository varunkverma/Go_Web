package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]

	template := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<h1>` + name + `</h1>
	</body>
	</html>`

	file, err := os.Create("index.html")
	defer file.Close()
	if err != nil {
		log.Fatalln("Error creating file: ", err)
	}
	io.Copy(file, strings.NewReader(template))
}
