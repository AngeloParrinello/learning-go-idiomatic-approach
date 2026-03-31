package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
)

// When a function executes
// as expected, nil is returned for the error parameter. If something goes wrong, an
// error value is returned instead.

// The error handling is indented inside
// an if statement. The business logic is not. This gives a quick visual clue to which
// code is along the “golden path” and which code is the exceptional condition.

// Two kinds of errors in Go: Simple Errors and Sentinel Errors

// Simple errors (just a plain string wrapped in an error type)
func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}
func doubleEvenWithErrorf(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

// Sentinel errors: (errors that starts with Err) if you have an error condition that indicates a specific state has been reached in
// your application where no further processing is possible and no contextual information needs
// to be used to explain the error state, a sentinel error is the correct choice.
// Sentinel
// errors should be rare, so they can be handled by convention instead of language rules.
// Yes, they are public package-level variables. This makes them mutable, but it’s highly
// unlikely someone would accidentally reassign a public variable in a package. In short,
// it’s a corner case that is handled by other features and patterns. The Go philosophy is
// that it’s better to keep the language simple and trust the developers and tooling than it
// is to add additional features.
func testErrFormat() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}
}
