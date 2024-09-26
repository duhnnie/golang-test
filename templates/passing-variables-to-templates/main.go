package main

import (
	"html/template"
	"os"
)

type Author struct {
	FirstName string
	LastName string
}

type Params struct {
	Author Author
}

type MyStruct struct {
	Params Params
}

func main() {
	myStruct := MyStruct{
		Params: Params{
			Author: Author{
				FirstName: "John",
				LastName: "Doe",
			},
		},
	}

	footerTplString := `
	{{define "footer"}}
		<footer>{{ .Params.Author.FirstName }}</footer>
	{{end}}
	`

	headerTplString := `
	{{define "header"}}
		Welcome to {{.FirstName}}'s blog!
	{{end}}
	`

	indexTplString := `
	{{template "header" .Params.Author}}
	<article>Here goes the content</article>
	{{template "footer" .}}
	`

	// template.New("header").Parse(headerTplString)
	// template.New("footer").Parse(footerTplString)
	tpl, _ := template.New("index.html").Parse(indexTplString)
	tpl.New("header").Parse(headerTplString)
	tpl.New("footer").Parse(footerTplString)

	tpl.Execute(os.Stdout, myStruct)
}
