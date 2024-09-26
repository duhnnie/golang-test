package main

import (
	"html/template"
	"os"
)

type Person struct {
	FirstName string
	LastName string
}

func main() {
	person := Person {
		FirstName: "Bobby",
		LastName: "Brown",
	}

	definitionTplString := `
	{{define "greet"}}
		{{ $firstname := "Bobby" }}
		<div>Hello, my name is {{.FirstName}}.<div>
		{{ $firstname = "John" }}
		<div>Actually, my name is {{ $firstname }}</div>
	{{end}}
	`

	tplString := `{{template "greet" .}}`

	tpl, _ := template.New("myTemplate").Parse(tplString)
	tpl.New("definition").Parse(definitionTplString)
	tpl.Execute(os.Stdout, person)
}
