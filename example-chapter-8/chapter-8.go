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

// Errors are value too and we can define custom ones
type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (e StatusErr) Error() string {
	return e.Message
}

// we always return error even though we return custom errors
//func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
//	err := login(uid, pwd)
//	if err != nil {
//		return nil, StatusErr{
//			Status:  InvalidLogin,
//			Message: fmt.Sprintf("invalid credentials for user %s", uid),
//		}
//	}
//	data, err := getData(file)
//	if err != nil {
//		return nil, StatusErr{
//			Status:  NotFound,
//			Message: fmt.Sprintf("file %s not found", file),
//		}
//	}
//	return data, nil
//}

// When using custom errors, never define a variable to be of the type
// of your custom error. Either explicitly return nil when no error
// occurs or define the variable to be of type error.

// When an error is passed back through your code, you often want to add additional
// context to it. This context can be the name of the function that received the error or
// the operation it was trying to perform. When you preserve an error while adding
// additional information, it is called wrapping the error. When you have a series of
//  wrapped errors, it is called an error chain.
