package main

import "fmt"

func stringp(s string) *string {
	return &s
}

func main() {
	x := "hello"
	pointerToX := &x // & is the address operator, it gives us the memory address of x
	fmt.Println("pointerToX:", pointerToX)
	fmt.Println("value at pointerToX:", *pointerToX) // * is the dereference/indirection operator, it gives us the value at the memory address

	y := *pointerToX + " world"
	fmt.Println("y:", y)

	z := 10
	var pointerToZ *int // pointer type is *int, it can hold the memory address of an int variable
	pointerToZ = &z
	fmt.Println("pointerToZ:", pointerToZ)

	var q = new(int)      // new allocates memory for an int and returns a pointer to it
	fmt.Println(q == nil) // false
	fmt.Println(*q)       // 0, new initializes the memory to the zero value of the type

	type person struct {
		name    string
		surname *string
		age     int
	}

	// This code does not compile because the struct literal is missing the surname field, which is a pointer to a string.
	// p := person{name: "Alice", surname: "Smith", age: 30}

	// We can bypass the issue by:
	// - Using a pointer to a string for the surname field
	// - Create a helper function that takes in a boolean, numeric, string (we can't use & before a primitive literal because they do not have memory address they exist only at compile time)
	//   and returns a pointer to that value
	p := person{name: "Alice", surname: stringp("Smith"), age: 30}
	fmt.Println(p)
}
