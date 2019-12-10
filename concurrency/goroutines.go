package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func timer() {
	time.Sleep(time.Second * 3)
}

func testmain() {
	say("Sync") // sync 하게 실행
	/*
		go say() Async 호출은 별도의 goroutine에서 동작한다.
		main routine은 계속 다음 문장 (time.Sleep) 을 실행한다.
		goroutine은 그 순서가 일정하지 않으므로 프로그램 실행시마다 다른 출력 결과를 나타낼 수 있다.
	*/
	go say("Async1") // Async
	go say("Async2") // Async
	go say("Async3") // Async

	timer()
	fmt.Println("Main Finished")

	var wait sync.WaitGroup
	wait.Add(2)

	// 익명함수 goroutine
	// go keyword 뒤에 익명함수를 바로 정의함, 이는 익명함수를 비동기로 실행

}
