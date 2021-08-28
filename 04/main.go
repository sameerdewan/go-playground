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
	values := map[string]string{
		"string1": "valueAtStringOne",
		"string2": "valueAtStringTwo",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", values)
	if err != nil {
		log.Fatalln(err)
	}
}
