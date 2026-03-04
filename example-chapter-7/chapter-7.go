package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Simple function with a receiver of type Person, this is a method of the Person type
// It is nonidiomatic to use this or self as the receiver name, it is more common to use a short abbreviation of the type name, in this case p for Person
func (p Person) String() string {
	return fmt.Sprintf("%s %s is %d years old", p.FirstName, p.LastName, p.Age)
}

func main() {

	p := Person{FirstName: "Alice", LastName: "Smith", Age: 30}
	fmt.Println(p.String()) // Output: Alice Smith is 30 years old

}
