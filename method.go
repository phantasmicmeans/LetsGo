package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

/*
 Method Receiver 메소드 리시버,
 메소드 리시버는 func 키워드와 메소드 이름 사이에 인자로 들어간다.
*/
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) basicAbs() float64 {
	return math.Sqrt(v.X*v.Y + v.X*v.Y)
}

/*
 기본타입들에도 메소들르 붙일 수 있다.
*/
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func testMain() {
	a := Vertex{3, 4}
	fmt.Println(v.Abs())

	v := &Vertex{3, 4}
	result := v.Abs()
	fmt.Println("result = ", result)

	f := MyFloat(100)
	fmt.Println(f.Abs())
}
