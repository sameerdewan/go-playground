package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"firstThreeLetters": firstThreeLetters,
	"uppercase":         strings.ToUpper,
}

func firstThreeLetters(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templates/index.gohtml"))
}

func main() {
	data := map[string]string{
		"Sameer": "Dewan",
		"Thomas": "Pain",
		"Martin": "Luther",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
