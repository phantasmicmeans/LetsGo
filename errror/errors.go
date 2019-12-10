package main

import (
	"fmt"
	"time"
)

// As with fmt.Stringer, the fmt package looks for the error interface when printing values.
// Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
