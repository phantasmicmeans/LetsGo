package main

import "fmt"

/*
	구조체
*/
type Vertex struct {
	X int
	Y int
}

func structTest() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
	fmt.Println(v.X)
}

/*
Pointer, Go에는 Pointer가 있지만 포인터 연산은 불가함.
구조체 변수는 구조체 포인터를 이용해 접근 가능
*/

func pointerStructTest() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p)
}

/*
 구조체 리터럴 (Struct Literals)
*/
func structLiterals() {
	var (
		p = Vertex{1, 2}  // has type Vertex
		q = &Vertex{1, 2} // has type *Vertex, 구조체 리터럴에 대한 포인터
		r = Vertex{X: 1}  // Y:0 is implict
		s = Vertex{}      // X:0 and Y:0
	)

	fmt.Println(p, q, r, s)
}

func structLiterals2() {
	var (
		p = Vertex{1, 2}
		q = &Vertex{1, 2}
		r = Vertex{X: 1}
		s = Vertex{}
	)
}

func testMain() {
	structTest()
	pointerStructTest()
	structLiterals()
}
