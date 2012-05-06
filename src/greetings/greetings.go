package greetings

import (
	"fmt"
)

type Greeting struct {
	Geschlecht int
	Name string
}

func Hello(people ...Greeting) {
	
	var anrede string

	for i, g := range people {
	
		switch g.Geschlecht {
		default:
			anrede = ""
		case 1:
			anrede = "Frau "
		case 2:
			anrede = "Herr "
		}
	
		greeting := "(%v) Hallo, %v%v!\n"
		
		fmt.Printf(greeting, i, anrede, g.Name)
	}
}
