package main

import (
	"os"
	"text/template"
)

func main() {
	tplString := `{{if .}} it's true {{else}} it's false {{end}}`
	tpl, _ := template.New("anotherTemplate").Parse(tplString)

	tpl.Execute(os.Stdout, false)

	tplString2 := "{{range .}} {{.}} {{else}} No elements in array {{end}}"
	tpl2, _ := template.New("template2").Parse(tplString2)

	tpl2.Execute(os.Stdout, []int{2, 4, 6, 8, 0})
	tpl2.Execute(os.Stdout, []int{})
}