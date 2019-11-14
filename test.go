package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/* Variable */

/*
 변수는 Go 키워드 var를 사용해 선언한다.
 var 키워드 뒤에 변수명을 적고 그 뒤에 변수타입을 적는다.
*/
func varTest() {
	var a int
	var f float32 = 11. // float32 타입의 변수 f에 11.0 이라는 초기값을 할당한다.

	a = 10
	f = 12.0

}
