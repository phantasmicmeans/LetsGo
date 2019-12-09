package main

import "fmt"

/*
The Empty Interface

The interface type that specifies zero methods is known as the empty interface.package LetsGo

 interface {}

An empty interface may hold values of any type.
(Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type.
For example, fmt.Println takes any number of arguments of type interfaces{}.
*/

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
