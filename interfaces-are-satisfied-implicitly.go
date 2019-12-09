package main

import "fmt"

/* Interfaces are implemented implicitly
- A type implements an interface by implementing its methods.
- There is no explict declaration of intent, no "implements" keyword.
*/
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// But we don't need to explicitly declare that it dose so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
