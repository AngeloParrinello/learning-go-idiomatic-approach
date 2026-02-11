package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

// Variadic input parameters and slices
// the variadic parameter must be the last parameter in the function signature
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// We can return multiple values from a function
// If the function can fail, it's common to return an error as the last return value
// we return nil if there is no error, or an error value if there is an error
func divAndRemainder(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, fmt.Errorf("denominator cannot be zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

// A common pattern is to return a cleanup function that the caller can defer to ensure that resources are properly released
func getFile(filename string) (*os.File, func(), error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, nil
}

// Functions are values too
func add(i, j int) int { return i + j }
func sub(i, j int) int { return i - j }
func mul(i, j int) int { return i * j }
func div(i, j int) int { return i / j }

// We can define a function type declaration and use it to declare variables
type opFuncType func(int, int) int

var opMap2 = map[string]opFuncType{}

// We can also return functions from other functions
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(10, a...))
	fmt.Println(addTo(10, []int{1, 2, 3}...))

	// ---------
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

	// ---------
	var opMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		[]string{"10", "+", "5"},
		[]string{"20", "-", "4"},
		[]string{"3", "*", "7"},
		[]string{"16", "/", "2"},
		[]string{"two", "+", "zero"},
		[]string{"5"},
	}
	for _, expr := range expressions {
		if len(expr) != 3 {
			fmt.Println("invalid expression:", expr)
			continue
		}
		// Convert the operands from string to int
		p1, err := strconv.Atoi(expr[0])
		if err != nil {
			fmt.Println("invalid operand:", expr[0])
			continue
		}

		op := expr[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("invalid operator:", op)
			continue
		}

		p2, err := strconv.Atoi(expr[2])
		if err != nil {
			fmt.Println("invalid operand:", expr[2])
			continue
		}

		result := opFunc(p1, p2)
		fmt.Printf("%d %s %d = %d\n", p1, op, p2, result)
	}

	// Anonymous functions
	// Typically used in combination with goroutines or defer statements
	for i := 0; i < 5; i++ {
		func(n int) {
			fmt.Println("printing from anonymous function:", n)
		}(i)
	}

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Alice", "Smith", 30},
		{"Bob", "Johnson", 25},
		{"Charlie", "Brown", 35},
	}
	fmt.Println(people)

	// Passing functions as arguments to other functions
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	// The people slice is changed by the call to sort.Slice
	fmt.Println(people)

	// We can also return functions from other functions
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}

	// Defer: programs often create temporary resources that need to be cleaned up, such as open files or network connections.
	// The defer statement allows you to schedule a function call to be run after the function completes,
	// which is useful for ensuring that resources are properly released.
	// The deferred function calls are executed in LIFO order (last-in, first-out) when the surrounding function returns.
	// This means that if you have multiple defer statements, they will be executed in reverse order of their appearance in the code.

	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// we need to close it after we use, no matter how we exit the function, even if there is an error
	// THIS RUNS AFTER THE RETURN STATEMENT
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}

	f, closer, err := getFile(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer closer()
}
