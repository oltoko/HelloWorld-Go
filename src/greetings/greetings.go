package greetings

import (
	"fmt"
)

func Hello(name string) {

	greeting := "Hello, %v!\n"
	
	fmt.Printf(greeting, name)
}

func Cheerio(name string) {

	greeting := "Cheerio, %v!\n"
	
	fmt.Printf(greeting, name)
}
