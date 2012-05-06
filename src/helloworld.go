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
	
	var blub32 float32 = 1.23
	var blub64 float64 = 1.23
	
	Check(nil)
	Check(123)
	Check(blub32)
	Check(blub64)
	Check(gC)
	Check("wtf")
}

