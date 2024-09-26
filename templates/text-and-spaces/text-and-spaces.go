package main

import (
	"os"
	"text/template"
)

func main() {
	tplString := "{{3 -}}   <  {{- 4}}"
	tpl, _ := template.New("templateName").Parse(tplString)

	tpl.Execute(os.Stdout, nil)
}