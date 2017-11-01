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

//Create struct containing player weapon types
type weapon struct {
	Name   string
	Weild  string
	Class  string
	Damage int
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
		Definition: "Swift and Adept a rogue uses cunningness to outwit his opponants.",
	}
	marksman := class{
		Name:       "Marksman",
		Definition: "Calculated and Precise a Marksman ambushes foes from above and afar.",
	}
	//Create the different weapons
	shortsword := weapon{
		Name:   "Short Sword",
		Weild:  "One Handed, Main",
		Class:  "Warrior, Rogue, Mage",
		Damage: 10,
	}
	twohandedsword := weapon{
		Name:   "Two-Handed Sword",
		Weild:  "Two Handed, Both",
		Class:  "Warrior",
		Damage: 25,
	}
	staff := weapon{
		Name:   "Magic Staff",
		Weild:  "Two Handed, Both",
		Class:  "Mage",
		Damage: 12,
	}
	longbow := weapon{
		Name:   "Longbow",
		Weild:  "Two Handed, Both",
		Class:  "Marksman",
		Damage: 14,
	}
	dagger := weapon{
		Name:   "Dagger",
		Weild:  "One Handed, Main or Offhand",
		Class:  "Rogue",
		Damage: 6,
	}
	//Create slices of class struct
	classes := []class{warrior, mage, rogue, marksman}
	weapons := []weapon{shortsword, twohandedsword, staff, longbow, dagger}
	//Combine the slices into a struct
	combined := struct {
		PlayerClasses []class
		PlayerWeapons []weapon
	}{
		PlayerClasses: classes,
		PlayerWeapons: weapons,
	}
	//Parse combined data
	err := tpl.ExecuteTemplate(os.Stdout, "01.gohtml", combined)
	if err != nil {
		log.Fatalln(err)
	}
}
