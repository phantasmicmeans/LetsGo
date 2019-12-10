package main

import (
	"fmt"
	"os"
)

/*
 1. 지연실행 defer
 defer keyword는 특정 문장 혹은 function을 defer를 호출하는 method가 리턴하기 직전에 실행하게 한다.
 일반적으로 defer는 java의 finally block처럼 마지막에 clean-up 작업을 위해 사용된다.

 2. panic 함수
 panic()은 현재 function을 즉시 멈추고, 현재 function에 defer 함수들을 모두 실행한 후 즉시 리턴한다.
 이러한 panic 모드 실행 방식은 다시 상위 함수에도 똑같이 적용된다.
 계속 call stack을 타고 올라가며 적용된다.
 그리고 마지막에는 program이 error를 내고 종료하게 된다.

 3. recover 함수
 recover()은 panic()에 의한 패닉상태를 다시 정상상태로 되돌리는 함수이다.
 위의 panic 예제에서는 main 함수에서 println()이 호출되지 못하고 crash되지만, recover()를 사용하면
 panic 상태를 제거하고 다음 문장 호출 가능
*/

func defferTest() {
	f, err := os.Open("C:/Users/SangMin/Desktop/golang/LetsGo/etc/test.txt")
	if err != nil {
		panic(err)
	}

	// main() 마지막에 file close 실행
	defer f.Close()

	// file 읽기
	bytes := make([]byte, 1024)
	f.Read(bytes)
	println("Content : ", bytes)
	println("length :", len(bytes))
}

func panicTest() {
	openFile("Invalid.txt")
	println("Done") // 이 문장 실행 안됨
}

func panicTestWithRecover() {
	openFileWithRecover("Invalid.txt")
	println("Done") // 이 문장 실행 안됨
}

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err) // defer 를 모두 실행시키고 종료, 상위 스택까지
	}

	defer f.Close()
}

func openFileWithRecover(fn string) {
	defer func() { // defer function, execute after panic()
		if r := recover(); r != nil { // r에 의해 panic 상태 제거
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}

func main() {
	var number int = 1
	defferTest()
	if number == 0 {
		panicTest()
	} else {
		panicTestWithRecover() // println("Done") execute successfully
	}
}
