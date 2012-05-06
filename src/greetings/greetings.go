package greetings

import (
	"fmt"
	"time"
)

const (
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
		Hello(value)
	default:
		fmt.Println("Unbekannter Datentyp :'(")
	}
}

type Gender int

type PersonTable map[string]*Person

func NewPersonTable() PersonTable {
	return make(PersonTable)
}

type Person struct {
	id        string
	gender    Gender
	firstName string
	lastName  string
	birthDate time.Time
	buddies   PersonTable
}

func NewPerson(id string) *Person {

	if len(id) == 0 {
		return nil
	}

	person := new(Person)

	person.id = id
	person.buddies = NewPersonTable()

	return person
}

func (p *Person) Hello(buddy bool) {

	if buddy {
		fmt.Printf("Hi %v\n", p.firstName)
	} else {
		s := p.salutation()
		fmt.Printf("Guten Tag %v %v", s, p.lastName)
	}
}

func (p *Person) salutation() string {
	switch p.gender {
	case Weiblich:
		return "Frau"
	case Maennlich:
		return "Herr"
	}

	return ""
}

func (p *Person) SetGender(gender Gender) {
	p.gender = gender
}

func (p *Person) Gender() Gender {
	return p.gender
}

func (p *Person) SetFirstName(firstName string) {
	p.firstName = firstName
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetLastName(lastName string) {
	p.lastName = lastName
}

func (p *Person) LastName() string {
	return p.lastName
}

func (p *Person) SetBirthDate(birthDate time.Time) {
	p.birthDate = birthDate
}

func (p *Person) BirthDate() time.Time {
	return p.birthDate
}
