package main

import (
	"fmt"
	"sync"
)

func main() {
	// Wait Group 생성, 2개의 goroutine을 기다림.
	// sync.WaitGroup은 여러 goroutine이 끝날 때까지 기다리는 역할.
	var wait sync.WaitGroup
	wait.Add(2) // 2개의 goroutine 추가

	// 익명함수를 사용한 goroutine
	go func() {
		defer wait.Done() // 끝나면 .Done() 호출
		fmt.Println("Hello")
	}()

	// 익명함수에 파라미터 전달
	go func(msg string) {
		defer wait.Done() // 끝나면 .Done() 호출
		fmt.Println(msg)
	}("Hi") // parameter 전달.

	// goroutine이 모두 끝날 때까지 대기
	wait.Wait()
}
