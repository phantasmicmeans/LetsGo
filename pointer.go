package main

import "fmt"

type Vertex struct {
	X, Y int
}

/*
 new(T)는 모든 필드가 0(zero value)이 할당된 T 타입의 포인터를 반환한다.
*/
func newTest() {
	var v *Vertex = new(Vertex)
	x := new(Vertex)
	fmt.Println(v)
	fmt.Println(x)
	// 같은 표현
}

func main() {
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)

	newTest()
}
