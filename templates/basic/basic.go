package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count uint
}

func main() {
	sweaters := Inventory{"wool", 17}
	tmplStr := "{{.Count}} items are made of {{.Material}}"
	tmpl, err := template.New("templateName").Parse(tmplStr)

	if err != nil { panic(err) }

	err = tmpl.Execute(os.Stdout, sweaters)

	if err != nil { panic(err) }
}