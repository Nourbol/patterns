package newspaper

import "fmt"

type Subscriber interface {
	HandleEvent(string)
}

type Person struct {
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		name,
	}
}

func (p *Person) HandleEvent(message string) {
	fmt.Printf("Dear %s, %s\n", p.Name, message)
}
