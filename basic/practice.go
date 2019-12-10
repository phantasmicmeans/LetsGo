package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

/*
if 에서 조건문 앞에 문장 실행
조건 비교 전, 'v := math.Pow(x, n)'을 실행함.
*/
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		fmt.Println("V : ", v)
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here
	return lim
}

func testMain() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
