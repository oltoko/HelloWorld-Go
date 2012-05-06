package greetings

import (
	"fmt"
)

func Hello(geschlecht int, name string) {
	
	var anrede string

	switch geschlecht {
	default:
		anrede = ""
	case 1:
		anrede = "Frau "
	case 2:
		anrede = "Herr "
	}

	greeting := "Hello, %v%v!\n"
	
	fmt.Printf(greeting, anrede, name)
}

func Cheerio(name string) {

	greeting := "Cheerio, %v!\n"
	
	fmt.Printf(greeting, name)
}
