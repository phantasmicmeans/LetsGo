package main

import "fmt"

/*
 함수 클로져 (Function Closures)
 adder function은 클로져(closure)를 반환합니다.
 각각의 클로져는 자신만의 sum 변수를 가집니다.
*/
func adder() func(x int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func testMain() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
