package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("vim-go")
	ifTest()
	mathTest()
	stringTest()
}

/* Variable */

/*
 변수는 Go 키워드 var를 사용해 선언한다.
 var 키워드 뒤에 변수명을 적고 그 뒤에 변수타입을 적는다.
*/
func varTest() {
	/*
		var a int
		var f float32 = 11. // float32 타입의 변수 f에 11.0 이라는 초기값을 할당한다.

		a = 10
		f = 12.0

		// 복수 변수들이 선언된 상황에서 초기값 지정. 초기값은 순서대로 변수에 할당.
		var i, j, k int = 1, 2, 3

		// Go 에서는 할당되는 값을 보고 그 타입을 추론하는 기능이 자주 사용된다.
		// 아래 코드에서 i는 정수형으로 1이 할당되고, s는 문자열로 Hi 가 할당된다.
		var h = 1
		var s = "HI"

		// var i = 1 을 쓰는 대신 i:=1이라고 var를 생략할 수 있다.
		// 이러한 표현은 함수 내에서만 가능하며, 함수 밖에서는 var를 사용해야한다. Go에서 변수와 상수는 함수밖에서도 사용할 수 있다.
	*/
}

func constTest() {
	const c int = 10
	const s string = "Hi"

	// 타입 추론
	const c1 = 10
	const s2 = "Hi"

	// 여러 상수들을 묶어서 지정 가능
	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)

	// iota identifier
	const (
		Apple  = iota // 0
		Grape         // 1
		Orange        // 2
	)
}

func dataType() {
	/* 불리언, 문자열, 정수형, Float 및 복소수 타입, 기타타입
	 1. bool
	 2. string - string은 수정될 수 없는 Immutable 타입
	 3. 정수형 타입
		- int int8 int16 int64
		- uint uint8 uint16 uint32 uint64 uintptr
	 4. Float 및 복소수 타입
		- float32 float64 complex64 complex127
	 5. 기타 타입
		- byte: uint8과 동일하며 바이트 코드에서 사용
		- rune: int32와 동일하며 유니코드 코드포인트에서 사용
	*/
}

func stringLiteral() {
	/**
	문자열 리터럴을 Back Quote('') 혹은 이중인용부호("")를 사용하여 표현할 수 있다.
	01. Back Quote('')로 둘러 싸인 문자열은 Raw String Literal이라 부른다.
	이 안에 문자열은 별도로 해석되지 않고 Raw String을 그대로 가진다.

	이중인용부호("")로 둘러 쌓인 문자열은 Interpreted String Literal이라 한다.
	Escape 문자열들은 특별한 의미로 해석된다.
	*/
	interLiteral := "hello\n\nhowaryou"
	fmt.Println(interLiteral)
}

func ifTest() {
	var k int = 1
	if k == 1 {
		println("One")
	} else if k == 2 {
		println("Two")
	} else {
		println("Other")
	}

	var max int = 200
	var i int = 1
	if val := i * 2; val < max {
		fmt.Println(i)
	}
}

func switchTest() {
	var name string
	var category int = 1

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)
}

func switchTest2() {
	score := 99
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}
}

func mathTest() {
	println("Happy", math.Pi, "Day")
	fmt.Printf("Now you have %g problems \n", math.Nextafter(2, 3))

	fmt.Println(addTest(42, 13))
	fmt.Println(addTest2(42, 13))
}

func stringTest() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func addTest(x int, y int) int {
	return x + y
}

func addTest2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
