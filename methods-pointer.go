package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/**
 You can declare methods with pointer receivers
**/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/* Pointer Receiver */
func (v *Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/* Pointer Argument */
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	var v Vertex = Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	// function with Pointer Argument 는 must take a pointer
	ScaleFunc(v, 5) // Compile Error
	//cannot use v (type Vertex) as type *Vertex in argument to ScaleFun
	ScaleFunc(&v, 5) // OK

	// methods with pointer receivers는 value와 pointer 둘다 가능
	v.Scale2(5) // value -> OK
	p := &v
	p.Scale2(10) // pointer -> OK
}
