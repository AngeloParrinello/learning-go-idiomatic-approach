// Package money: (Package-level comment) package money provides various utilities to make it easy to manage money.
package money

import "fmt"

// Money represents the combination of an amount of money
// and the currency the money is in.
type Money struct {
	Value    float64
	Currency string
}

// Convert converts the value of one currency to another.
//
// It has two parameters: a Money instance with the value to convert,
// and a string that represents the currency to convert to. Convert returns
// the converted currency and any errors encountered from unknown or unconvertible
// currencies.
// If an error is returned, the Money instance is set to the zero value.
//
// Supported currencies are:
// USD - US Dollar
// CAD - Canadian Dollar
// EUR - Euro
// INR - Indian Rupee
//
// More information on exchange rates can be found
// at https://www.investopedia.com/terms/e/exchangerate.asp
func Convert(from Money, to string) (Money, error) {
	// ...
	return from, nil
}

// Go includes a command-line tool called go doc that views godocs. The command go
// doc PACKAGE_NAME displays the package godocs for the specified package and a list of
// the identifiers in the package. Use go doc PACKAGE_NAME.IDENTIFIER_NAME to display
// the documentation for a specific identifier in the package.

// Go allows blank imports, where the name assigned to an import is the underscore (_). Just as an
// underscore allows you to skip an unused return value from a function, a blank import
// triggers the init function in a package but doesn’t give you access to any of the
// exported identifiers in the package:
// import (
//  "database/sql"
//  _ "github.com/lib/pq"
// )
// This pattern is considered obsolete because it’s unclear that a registration operation is
// being performed.

// Go does not allow you to have a circular dependency between packages

// When you want to rename or move an exported type, you have to use an alias. Quite
// simply, an alias is a new name for a type.

type Foo struct {
	x int
	S string
}

func (f Foo) Hello() string {
	return "hello"
}
func (f Foo) goodbye() string {
	return "goodbye"
}

type Bar = Foo

// The alias can even be assigned to a variable of the original type without a type conversion
func MakeBar() Bar {
	bar := Bar{
		// these field names cannot be changed!
		x: 20,
		S: "Hello",
	}
	var f Foo = bar
	fmt.Println(f.Hello())
	return bar
}

// If you want to add new methods or change the fields in an aliased struct, you must add
// them to the original type.
