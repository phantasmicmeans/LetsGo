package main

/*
Interface Values

interface values can be thought of as a tuple of a value and a concrete type:
  (value, type)

An interfaces value holds a vlue of a specific underlying concrete type.
Calling a method on an interface value executes the method of the
same name on its underlying type.
*/

import (
	"fmt"
	"math"
)

type I interface {
	M()
	M2()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println("(t *T) M() methods called")
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println("(f F) M() methods called")
	fmt.Println(f)
}

func (t *T) M2() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func (f F) M2() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	i = &T{"Hello M()"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	/* Interface values with nil underlying values
	If the concrete value inside the interface itself is nil.
	the method will be called with a nil receiver.

	In some languages this would trigger a null pointer exception,
	but in Go it is common to wirte methods that gracefully handle being called
	with a nil receiver.

	Note that an interface value that holds a nil concrete value is itself non-nil.
	*/
	var i2 I
	i2 = &T{"Hello M2()"}
	describe(i2)
	i2.M2()
}
