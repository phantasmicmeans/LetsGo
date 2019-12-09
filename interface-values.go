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

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}
