package greetings

import "fmt"

const(
	Unbekannt = iota
	Weiblich
	Maennlich
)

type Greeting struct {
	Geschlecht int
	Name       string
}

func Hello(people ...Greeting) {

	var anrede string

	for _, g := range people {

		switch g.Geschlecht {
		default:
			anrede = ""
		case Weiblich:
			anrede = "Frau "
		case Maennlich:
			anrede = "Herr "
		}

		greeting := "Hallo, %v%v!\n"

		fmt.Printf(greeting, anrede, g.Name)
	}
}
