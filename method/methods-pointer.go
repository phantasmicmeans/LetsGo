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

func (v *Vertex) AbsP() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/**
 You can declare methods with pointer receivers
**/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/* Value Receiver -> Value / Pointer 둘다 가능 */
func (v Vertex) ScaleV(f float64) float64 {
	v.X = v.X * f
	v.Y = v.Y * f
	return v.X
}

func ScaleFunc2(v Vertex, f float64) float64 {
	v.X = v.X * f
	v.Y = v.Y * f
	return v.X
}

/* Pointer Receiver -> Value / Pointer 둘다 가능 */
func (v *Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/* Pointer Argument -> Pointer 만 가능 */
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/**
Choosing a value or pointer receiver
- Pointer Receiver 를 사용해야하는 2가지 이유
1. 첫번째는 method가 receiver points를 통해 value를 변경할 수 있다.
2. 두번째는 each method call 마다 value를 copy하지 않아도 된다.
**/
func testMain() {
	var v Vertex = Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	// function with Pointer Argument 는 must take a pointer
	//ScaleFunc(v, 5) // Compile Error
	//cannot use v (type Vertex) as type *Vertex in argument to ScaleFun
	ScaleFunc(&v, 5) // OK

	// methods with pointer receivers는 value와 pointer 둘다 가능
	v.Scale2(5) // value -> OK
	p := &v
	p.Scale2(10) // pointer -> OK

	var v2 Vertex = Vertex{3, 4}
	fmt.Println(v2.ScaleV(5))
	p2 := &v2
	fmt.Println(p2.ScaleV(5))

	fmt.Println(ScaleFunc2(v2, 5))
	fmt.Println(ScaleFunc2(*p2, 5))

	/* 일반적으로 모든 메소드는 둘다 value 타입이거나 둘다 pointer 타입을 사용한다. 두개를 함께 사용하지 않는다. */
	var v3 Vertex = Vertex{4, 5}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v3, v3.AbsP())
	v3.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v3, v3.AbsP())

}
