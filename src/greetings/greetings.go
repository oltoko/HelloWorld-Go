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

func Check(was interface{}) {
	
	switch value := was.(type) {
	case nil:
		fmt.Println("Du hast mir nichts gegeben -.-")
	case int:
		fmt.Printf("Dies ist ein integer mit dem Wert %v\n", value)
	case float64:
		fmt.Printf("Dies ist ein 64-bit float mit dem Wert %v\n", value)
	case Greeting:
		Hello(value);
	default:
		fmt.Println("Unbekannter Datentyp :'(")
	}
}