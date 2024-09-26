package main

import (
	"os"
	"text/template"
)

func main() {
	tplString := `{{/* a comment */}}
{{- /* a comment with white space trimmed from preceding and following text */ -}}
	A comment; discarded. May contain newlines.
	Comments do not nest and must start and end at the
	delimiters, as shown here.`

	tpl, _ := template.New("templateName").Parse(tplString)
	tpl.Execute(os.Stdout, nil)
}