package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 123
	}()

	var i int
	i = <-ch // 123 수신
	println(i)

	routineTest()
	bufferChannel()

	channelParam()
	channelSelect()
}

func routineTest() {
	fmt.Println("routineTest()")
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	// 위의 Go routine이 끝날 때까지 대기
	isTrue := <-done
	fmt.Printf("isTrue : %v \n\n", isTrue)
}

func bufferChannel() {
	fmt.Println("bufferChannel()")
	ch := make(chan int, 5)
	ch <- 101
	fmt.Println(<-ch)
}

func channelParam() {
	fmt.Println("channelParam()")
	ch := make(chan string, 2)
	sendChan(ch)
	receiveChan(ch)
}

/*
	Channel Parameter
	채널을 함수의 파라미터로 전달
	- 송신 parameter는 [p chan <- int]와 같이 chan <- 을 사용.
	- 수신 parameter는 [p <- chan int]와 같이 <- chan을 사용한다.
*/
func sendChan(ch chan<- string) {
	ch <- "Data"
	ch <- "Send"
	close(ch)
}

func receiveChan(ch <-chan string) {
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// <- ch는 두개의 리턴값을 가짐. 첫째는 채널 메시지, 두번째는 수신이 제대로 되었는지를 나타냄.
	// 채널이 제대로 close() 되었다면, 두번째 리턴값은 false를 리턴한다.
	if _, success := <-ch; !success {
		println("no more data")
	}
}

/*
	Channel Range문
	채널에서 송신자가 송신 후, 채널을 닫을 수 있다.
	그리고 수신자는 임의의 갯수의 데이터를 채널이 닫힐 때까지 계속 수신할 수 있다.
*/
func channelRange() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	// 방법 1, channel이 닫힌 것을 감지할 때까지 계속 수신.
	for {
		if i, success := <-ch; success {
			println(i)
		} else {
			break
		}
	}

	// 방법 2, 위 표현과 동일한 채널 range 문
	for i := range ch { // channel 이 닫히기 전 까지 계속 수신
		println(i)
	}
}

/*
Go의 select문은 복수 채널들을 기다리며 준비된(데이터를 보내온) 채널을 실행하는 기능을 제공한다.
즉, select 문은 여러개의 case문에서 각각 다른 채널을 기다리다가 준비가 된 channel case를 실행한다.
select문은 case channel들이 준비되지 않으면 계속 대기하고, 가장 먼저 도착한 채널의 case를 실행한다.

만약 복수 channel에서 신호가 오면, Go Runtime이 랜덤하게 그 중 한개를 선택한다.
하지만 select문에 default문이 있으면 case channel이 준비되지 않더라도 계속 대기하지 않고 바로 default문을 실행한다.
*/
func channelSelect() {
	fmt.Println("channelSelect()")
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 completed")
		case <-done2:
			println("run2 completed")
			break EXIT
		}
	}

	fmt.Println("EXIT!!")
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
