package main

import (
	"fmt"
)

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

	// Rather than declare that some variables and parameters are immutable, Go developers use pointers to indicate that a parameter is mutable.
	// if a pointer is passed to a function, the function gets a copy of the pointer.
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

	// Passing a map to a function means that you are copying a pointer to the map's underlying data structure
	// Becasue of this you shuld avoid using maps for input parameters or retrunv values

	// Passing a slice to a function has more complicated behavior: any modification to the contents of the slice is
	// reflected in the original variable, but using apopend to change the lenghtt isn't reflect in the original variable
	// Slices are freqauently passed around in Go programs
	// By default, you should assume that a slice is not modifiedd by a function. Your function's doc should specify if it modifies the slice's contents

	// Slices as buffer
	// Slices are often used as buffers for reading data from a file or network connection.

	//var filename string
	//file, err := os.Create(filename)
	//if err != nil {}
	//defer file.Close()
	//data := make([]byte, 100) // create a slice of bytes with a length of 100
	//for { // for loop to read data from the file until we reach the end of the file
	//	count, err := file.Read(data)
	//	if err != nil {}
	//	if count == 0 { // if count is 0, we have reached the end of the file
	//		break
	//	}
	//	// process the data read from the file
	//	fmt.Println("data read:", data[:count]) // print the data read from the file, we need to slice the data to only print the bytes that were read
	//}

}
