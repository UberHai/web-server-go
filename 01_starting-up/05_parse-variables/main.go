package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("srv/*.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "01.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "02.gohtml", "Shawn Wesley Stevens")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "03.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
