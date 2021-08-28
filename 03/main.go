package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	data := "This is the data!"
	err := tpl.ExecuteTemplate(os.Stdout, "tmp.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
