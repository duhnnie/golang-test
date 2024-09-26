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

	myMap := map[string]interface{}{
		"Params": map[string]interface{}{
			"Author": map[string]interface{} {
				"FirstName": "Bob",
				"LastName": "Mould",
			},
		},
	}

	tpl, _ := template.ParseFiles("index.html")

	// Here the template works fine with both struct and map,
	// cause they have the same shape
	tpl.Execute(os.Stdout, myStruct)
	tpl.Execute(os.Stdout, myMap)
}
