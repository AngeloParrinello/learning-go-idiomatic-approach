package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

// Go does not support inheritance, it encourages code reuse via built-in support for compostion and promotion:

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("Employee %s with ID %s", e.Name, e.ID)
}

type Manager struct {
	Employee   // this is an embedded field, it promotes the fields and methods of Employee to Manager
	Department string
}

func (m Manager) Description() string {
	// we can use the promoted fields and methods of Employee directly in Manager, we can also call the Description method of Employee to reuse its implementation
	fmt.Sprintf("Employee description: %s", m.Employee.Description())
	return fmt.Sprintf("Manager %s with ID %s, Department: %s", m.Name, m.ID, m.Department)
}

// but embedding is not inheritance!
// var eFail Employee = Manager{Employee: Employee{Name: "Alice", ID: "123"}, Department: "Sales"} // this does not compile because Manager is not an Employee, it is a different type that embeds Employee

// No dynamic dispatch in Gom the methods on the embedded field have no idea they are embedded.

// While Go’s concurrency model gets all the publicity, the real star of Go’s
//design is its implicit interfaces, the only abstract type in Go.

// interface definition
// Interfaces are usually named with “er” endings
// In an interface declaration, an interface literal appears after the name of the interface
// type. It lists the methods that must be implemented by a concrete type to meet the
// interface. The methods defined by an interface are called the method set of the
// interface.
type Stringer interface {
	String() string
}

// So far, nothing that’s been said is much different from interfaces in other languages.
// What makes Go’s interfaces special is that they are implemented implicitly.
//concrete type does not declare that it implements an interface.

// JS, Python, Ruby are duck typed languages (If it
//walks like a duck and quacks like a duck, it’s a duck)

// Interfaces specify what callers need. The client code defines the
// interface to specify what functionality it requires.
type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	// business logic
	return "processed data"
}

type Logic interface {
	Process(data string) string
}
type Client struct {
	L Logic // basically, Client needs to implement the method Process
}

func (c Client) Program() {
	// get data from somewhere
	data := "some data"
	c.L.Process(data)
}

// interfaces in Go encourage the Decorator pattern
// It is common in Go to write factory functions that take in an instance of an interface and
// return another type that implements the same interface.

//For example, say you have a
//function with the following definition:
//func process(r io.Reader) error
//You can process data from a file with the following code:
//r, err := os.Open(fileName)
//if err != nil {
//return err
//}
//defer r.Close()
//return process(r)
//return nil
//The os.File instance returned by os.Open meets the io.Reader interface and can be
//used in any code that reads in data. If the file is gzip-compressed, you can wrap the
//io.Reader in another io.Reader:
//r, err := os.Open(fileName)
//if err != nil {
//return err
//}
//defer r.Close()
//gz, err = gzip.NewReader(r)
//if err != nil {
//return err
//}
//defer gz.Close()
//return process(gz)
//Now the exact same code that was reading from an uncompressed file is reading from
//a compressed file instead.

// Rule of thumb, ACCEPT INTERFACE, RETURN STRUCTS

// When an interface could be one of multiple possible types, use a type switch instead:
func doThings(i interface{}) {
	// Instead of specifying a boolean operation, you specify a variable of an
	// interface type and follow it with .(type).
	// Since the purpose of a type switch is to derive a new variable from
	// an existing one, it is idiomatic to assign the variable being switched
	// on to a variable of the same name (i := i.(type)), making this
	// one of the few places where shadowing is a good idea. To make the
	// comments more readable, our example doesn’t use shadowing.
	switch j := i.(type) {
	case nil:
		fmt.Println(j)
	// i is nil, type of j is interface{}
	case int:
	// j is of type int
	case io.Reader:
	// j is of type io.Reader
	case string:
	// j is a string
	case bool, rune:
	// i is either a bool or rune, so j is of type interface{}
	default:
		// no idea what i is, so j is of type interface{}
	}
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

	var client = Client{
		L: LogicProvider{}, // here we did not use Logic but LogicProvider!
	}
	client.Program()

	var s *string
	fmt.Println(s == nil) // true
	var y interface{}
	fmt.Println(y == nil) // true
	y = s
	fmt.Println(y == nil) // false
	// In the Go runtime, interfaces are implemented as a pair of pointers, one to the under‐
	// lying type and one to the underlying value. As long as the type is non-nil, the inter‐
	// face is non-nil.

	// you need a way to say that a variable could
	// store a value of any type. Go uses interface{}
	// An empty interface type
	// simply states that the variable can store any value whose type implements zero or
	// more methods.
	var i interface{}
	i = 20
	i = "hello"
	i = struct {
		FirstName string
		LastName  string
	}{"Fred", "Fredson"}
	fmt.Println(i)

	// One common use of the empty interface is as a placeholder for data of uncertain
	// schema that’s read from an external source, like a JSON file:
	// one set of braces for the interface{} type,
	// the other to instantiate an instance of the map
	data := map[string]interface{}{}
	contents, err := os.ReadFile("testdata/sample.json")
	if err != nil {
		return
	}
	fmt.Println(string(contents))
	fmt.Println(data)
	json.Unmarshal(contents, &data)

	// If you see a function that takes in an empty interface, it’s likely that it is using reflection
	// to either populate or read the value.

	// These situations should be relatively rare. Avoid using interface{}. As we’ve seen,
	// Go is designed as a strongly typed language and attempts to work around this are
	// unidiomatic.

	// A type assertion names the concrete type that implemented the interface,
	// or names another interface that is also implemented by the concrete type underlying
	// the interface.
	type MyInt int
	var z interface{}
	var mine MyInt = 20
	z = mine
	i2, ok := z.(MyInt) // i2 is of type MyInt
	// comma ok Idiom
	if !ok {
		fmt.Println("not a MyInt")
	}
	fmt.Println(i2 + 1)
	// A type assertion is very different from a type conversion. Type conversions
	// can be applied to both concrete types and interfaces and
	// are checked at compilation time. Type assertions can only be
	// applied to interface types and are checked at runtime. Because they
	// are checked at runtime, they can fail. Conversions change,
	// assertions reveal.

}
