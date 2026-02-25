package main

import "fmt"

func stringp(s string) *string {
	return &s
}

func failedUpdate(g *int) {
	x := 10
	g = &x // this does not update the value of g in the caller, it only updates the local copy of g in this function
}

func update(g *int) {
	*g = 20 // this updates the value at the memory address that g points to, which is the same memory address that the caller's pointer points to
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

	// Rather than declare taht some variables and parameters are immutable, Go developers use pointers to indicate that a parameter is mutable.
	// if a pointer is passed to a function, the fucntion gets a copy of the pointer.
	// Two implications:
	// - when you pass a nil pointer to a function you cannot make the value non nil
	var f *int // f is nil
	failedUpdate(f)
	fmt.Println("f after failedUpdate:", f) // f is still nil
	// -if you want the value assigned to a pointer parameter to still bne there when you exit the function, you must dereference the pointer and set the value.
	a := 10
	failedUpdate(&a)
	fmt.Println("a after failedUpdate:", a) // a is still 10
	update(&a)
	fmt.Println("a after update:", a) // a is now 20
}
