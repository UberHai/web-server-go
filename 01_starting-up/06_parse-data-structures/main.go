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

	//Parse header
	err := tpl.ExecuteTemplate(os.Stdout, "01.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	//Parse body content
	//Create a slice of data to parse
	data := []string{"Shawn", "Wesley", "Stevens"}
	err = tpl.ExecuteTemplate(os.Stdout, "02a.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

	//Create a map of string:string to parse
	moredata := map[string]string{
		"first":  "He loves programming!",
		"second": "He enjoys gaming!",
		"third":  "He programs games!"}
	err = tpl.ExecuteTemplate(os.Stdout, "02b.gohtml", moredata)
	if err != nil {
		log.Fatalln(err)
	}

	//Create a struct to parse
	type langRatings struct {
		Name   string
		Rating string
	}
	nextdata := langRatings{
		"Golang Rated:",
		"4.9/5.0 Stars",
	}
	err = tpl.ExecuteTemplate(os.Stdout, "02c.gohtml", nextdata)
	if err != nil {
		log.Fatalln(err)
	}

	//Parse footer
	err = tpl.ExecuteTemplate(os.Stdout, "03.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
