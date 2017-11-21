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

//Create struct containing player class types
type class struct {
	Name       string
	Definition string
}

func main() {
	//Create the different classes
	warrior := class{
		Name:       "Warrior",
		Definition: "Strong and courageous. Warriors stand ground and never falter before their foes.",
	}
	mage := class{
		Name:       "Mage",
		Definition: "Intellect and Wisdom are compelling to a mysterious mind that is a mage.",
	}
	rogue := class{
		Name:       "Rogue",
		Definition: "Swift and Adept a rogue uses cunningness to outwit his opponents.",
	}
	marksman := class{
		Name:       "Marksman",
		Definition: "Calculated and Precise a Marksman ambushes foes from above and afar.",
	}
	//Create slices of class struct
	classes := []class{warrior, mage, rogue, marksman}
	//Parse classes data
	err := tpl.ExecuteTemplate(os.Stdout, "01.gohtml", classes)
	if err != nil {
		log.Fatalln(err)
	}

}
