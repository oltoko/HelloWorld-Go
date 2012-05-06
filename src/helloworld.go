package main 

import . "./greetings"

func main() {

	var (
		gA Greeting = Greeting{Weiblich, "Müller"}
		gB Greeting = Greeting{Maennlich, "Schmidt"}
		gC Greeting = Greeting{Unbekannt, "Welt"}
	)
	
	Hello()
	Hello(gA, gB)
	Hello(gC)
}

