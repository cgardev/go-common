package must

import (
	"log"
)

var errorAction = func(err error) {
	log.Fatal(err)
}

// SetErrorAction sets the action to be taken when an error is encountered.
func SetErrorAction(action func(err error)) {
	errorAction = action
}

// Do is a helper function that will log.Fatal if the error is not nil.
func Do[T any](r1 T, err error) T {
	if err != nil {
		errorAction(err)
	}
	return r1
}

// Do2 is a helper function that will log.Fatal if the error is not nil.
func Do2[T1 any, T2 any](r1 T1, r2 T2, err error) (T1, T2) {
	if err != nil {
		errorAction(err)
	}
	return r1, r2
}

// Do3 is a helper function that will log.Fatal if the error is not nil.
func Do3[T1 any, T2 any, T3 any](r1 T1, r2 T2, r3 T3, err error) (T1, T2, T3) {
	if err != nil {
		errorAction(err)
	}
	return r1, r2, r3
}

// DoVoid is a helper function that will log.Fatal if the error is not nil.
func DoVoid(err error) {
	if err != nil {
		errorAction(err)
	}
}
