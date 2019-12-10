package main

import (
	"bytes"
	"fmt"
	"sync"
)

type Vertex struct {
	X, Y int
}

type Person struct {
	name string
	age  int
}

/*
 SyncedBuffer를 구현하기 위한 자료형,
 new를 이용해 생성시, buffer는 빈 버퍼로 세팅, lock의 경우 Unlock 상태인 mutex로 됨. 즉시 사용 가능
*/
type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

/*
 new(T)는 모든 필드가 0(zero value)이 할당된 T 타입의 포인터를 반환한다.
 func new(Type) * Type
*/
func newTest() {
	var v *Vertex = new(Vertex)
	x := new(Vertex)
	fmt.Println(v)
	fmt.Println(x)
	// 같은 표현
}

func newTest2() {
	// new는 메모리 할당 단계에서 특정 값으로 초기화 시키는 것이 불가능하다.
	// 만약 특정 값으로 초기화를 함께 진행하고 싶은 경우는 composite literal 방식을 주로 사용한다.
	p1 := new(Person)
	p1.name = "sangmin"
	p1.age = 26

	p2 := &Person{"sangmin", 26}
	p3 := &Person{
		name: "sangmin",
		age:  26,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}

func testMain() {
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)

	newTest()
	newTest2()
}
