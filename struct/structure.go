package _struct

// Структура Person
type Person struct {
	Name   string
	Height int
	Weight int
}

func CreatePerson(name string, height, weight int) *Person {
	return &Person{name, height, weight}
}
