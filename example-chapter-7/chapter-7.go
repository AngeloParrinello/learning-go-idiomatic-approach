package main

import (
	"fmt"
	"time"
)

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

type Counter struct {
	total       int
	lastUpdated time.Time
}

// Pointer receiver
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// Value receiver
func (c Counter) String() string {
	return fmt.Sprintf("Counter total: %d, last updated: %s", c.total, c.lastUpdated.Format(time.RFC3339))
}

func main() {

	p := Person{FirstName: "Alice", LastName: "Smith", Age: 30}
	fmt.Println(p.String()) // Output: Alice Smith is 30 years old

	// Pointer receiver VS Value receiver
	// If your method modifies the receiver, you MUST use a pointer receiver
	// If your method needs to handle nil instances then it MUST use a pointer receiver
	// If your method does not modify the receiver, you CAN use a pointer receiver
	var c Counter
	fmt.Println(c.String())
	// notice here that we did not need to take the address of c to call the Increment method,
	//Go automatically takes the address of c when we call a method with a pointer receiver on a value type
	// under the hood, Go is doing this:
	// (&c).Increment()
	// but it's done for the local variable, not the passing of the variable to the method
	c.Increment()
	fmt.Println(c.String())

	// no get or set in Go structs, we can just access the fields directly

}
