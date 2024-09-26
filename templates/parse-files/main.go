package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	tpl, _ := template.ParseFiles("hello.html", "common.html")

	fmt.Printf("Current template: %s", tpl.Name())

	tpl.Execute(os.Stdout, map[string]string{
		"name": "Daniel",
	})

	tpl.ExecuteTemplate(os.Stdout, "footer", nil)

	// Set current template to "header"
	if tpl = tpl.Lookup("header"); tpl != nil {
		tpl.Execute(os.Stdout, nil)
	}
}