package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"os"
	"reflect"
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
	err     error
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

// use %w wraps at the end of the error format string and make the error to be
// wrapped the last parameter passed to fmt.Errorf.
func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

// If you want to wrap an error with your custom error type, your error type needs to
// implement the method Unwrap. This method takes in no parameters and returns an
// error.
func (se StatusErr) Unwrap() error {
	return se.err
}

// Wrapping errors is a useful way to get additional information about an error, but it
// introduces problems. If a sentinel error is wrapped, you cannot use == to check for it,
// nor can you use a type assertion or type switch to match a wrapped custom error. Go
// solves this problem with two functions in the errors package, Is and As.

// By default, errors.Is uses == to compare each wrapped error with the specified
// error. If this does not work for an error type that you define (for example, if your
// error is a noncomparable type), implement the Is method on your error:
type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		// It can compare anything, including slices
		return reflect.DeepEqual(me, me2)
	}
	return false
}

// The errors.As function allows you to check if a returned error (or any error it
// wraps) matches a specific type.

// err := AFunctionThatReturnsAnError()
// var myErr MyErr
// if errors.As(err, &myErr) {
//	fmt.Println(myErr.Code)
// }

// RULE OF THUMB: Use errors.Is when you are looking for a specific instance or specific
//values. Use errors.As when you are looking for a specific type.

// Go generates a panic whenever there is a situation where the
// Go runtime is unable to figure out what should happen next.
// As soon as a panic happens, the current function exits immediately
// and any defers attached to the current function start running. When those
// defers complete, the defers attached to the calling function run, and so on,
// until main is reached. The program then exits with a message and a stack trace.

// If there are situations in your programs that are unrecoverable, you can create your
// own panics.
func doPanic(msg string) {
	panic(msg)
}

// Go provides a way to capture a panic to provide a more graceful shutdown or to prevent
// shutdown at all. The built-in recover function is called from within a defer to check if
// a panic happened. If there was a panic, the value assigned to the panic is returned.
// Once a recover happens, execution continues normally.

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

// Reserve panics for fatal situations and use
// recover as a way to gracefully handle these situations.

// Idiomatic Go favors code that explicitly outlines the possible failure conditions
// over shorter code that handles anything while saying nothing.

// There is one situation where recover is recommended. If you are creating a library
// for third parties, do not let panics escape the boundaries of your public API. If a
// panic is possible, a public function should use a recover to convert the panic into an
// error, return it, and let the calling code decide what to do with them.

func main() {
	err := fileChecker("test.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	// You don’t usually call errors.Unwrap directly. Instead, you use
	//	// errors.Is and errors.As to find a specific wrapped error.
	//	if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
	//		fmt.Println(wrappedErr)
	//	}
	//}
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
		}
	}

	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}
