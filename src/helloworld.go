package main 

import . "./greetings"

func main() {

	var (
		gA Greeting = Greeting{1, "Müller"}
		gB Greeting = Greeting{2, "Schmidt"}
		gC Greeting = Greeting{99, "Welt"}
	)
	
	Hello()
	Hello(gA, gB)
	Hello(gC)
}

