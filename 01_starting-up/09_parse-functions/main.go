package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type class struct {
	Name       string
	Definition string
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("srv/01.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {

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
		Definition: "Swift and Adept a rogue uses cunningness to outwit his opponants.",
	}
	marksman := class{
		Name:       "Marksman",
		Definition: "Calculated and Precise a Marksman ambushes foes from above and afar.",
	}

	//Create a slice of class to parse
	moredata := []class{warrior, rogue, mage, marksman}
	err := tpl.ExecuteTemplate(os.Stdout, "01.gohtml", moredata)
	if err != nil {
		log.Fatalln(err)
	}

}
