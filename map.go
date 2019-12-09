package main

import "fmt"

/*
 맵은 값에 키를 지정.
 맵은 반드시 사용하기 전, make를 명시해야함 !!! (new가 아님)
 make를 수행하지 않은 nil에는 값을 할당할 수 없음.

 make는 new와 달리 해당 타입의 포인터값이 아닌, 값 그 자체를 반환한다.

*/
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var m2 map[string]Vertex

func mapTest2() {
	m2 = make(map[string]Vertex)
	m["hello"] = Vertex {
		40.342, -34.4324,
	}
	fmt.Println(m["sangmin"])
}
func mapTest() {
	m = make(map[string]Vertex)
	m["sangmin"] = Vertex{
		40.6342, -84.4323,
	}

	fmt.Println(m["sangmin"])
}

/*
 Map Literals, 구조체 리터럴과 비슷. Key를 반드시 지정해줘야함
*/
func mapLiterals() {
	var m = map[string]Vertex{
		"sangmin": Vertex{
			1, 3,
		},
		"hello": Vertex{
			3, 4,
		},
	}
	fmt.Println("m == ", m)

	var m2 = map[string]Vertex{
		"sangmin": {1, 3},
		"hello":   {3, 4},
	}
	fmt.Println("m2 == ", m2)
}

func mutatingMaps() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value : ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value : ", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value : ", m["Answer"])

	// key 존재여부 확인, 존재하면 ok == true, 존재하지 않으면 false.. element 타입에 따라 0(zero value)가 된다.
	v, ok := m["Answer"]
	fmt.Println("The value : ", v, " Present?", ok)
}

func testMain() {
	mapTest()
	fmt.Println()

	mapLiterals()
	fmt.Println()

	mutatingMaps()
}
