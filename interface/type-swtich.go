package main

import "fmt"

/*
Type assertions

A type assertion provides access to an interface value's
underlying concrete value.
  t := i.(T)
This statement asserts that the interface value i holds the concrete type T
and assigns the underlying T value to the variable t.

If i does not hold a T, the statment will trigger a panic.

To test whether an interface value holds a specific type,
a type assertion can return two values;
the underlying value and a boolean value that reports whether the assertion successded.
  t, ok := i.(T)
*/

func testmain() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	/*
		Type switches
		A type switch is a construct that permits several type assertions in series.

		A type switch is like a regular switch statement, but the cases in a type switch
		specify types (not values), and those values are compared against the type of the value
		held by the given interface value.

			switch v := i.(type) {
			case T:
				// here v has Type T
			case S:
				// here v has Type S
			default:
				// no match, here v has the same type as i
			}
	*/
	do(21)
	do("Hello")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
