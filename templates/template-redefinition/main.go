package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	baseTpl, _ := template.ParseFiles("base.html")
	baseTpl.Execute(os.Stdout, nil)

	fmt.Println("-------------------------")

	mainTpl, _ := template.ParseFiles("base.html", "index.html")
	mainTpl.Execute(os.Stdout, nil)

	fmt.Println("-------------------------")

	postTpl, _ := template.ParseFiles("base.html", "posts.html")
	postTpl.Execute(os.Stdout, nil)
}