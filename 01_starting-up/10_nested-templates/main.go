package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

//I have defined all templates inside templates/templates.gohtml
//I have called them from within index.gohtml
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 1337)
	if err != nil {
		log.Fatalln(err)
	}
}