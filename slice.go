package main

import "fmt"

/*
 Slice는 배열의 값을 가리킨다. (point), 그리고 배열의 길이를 가지고 있다.
 []T 는 타입 T를 가지는 요소의 slice이다.
*/
func sliceBasicTest() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p == ", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
	fmt.Println("\n")
}

/*
  슬라이스는 재분할 할 수도 있고, 같은 배열을 가리키는(point) 새로운 슬라이스를 만들 수도 있다.
*/
func slicingSlices() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p == ", p)

	/* index 1 ~ 3 */
	fmt.Println("p[1:4] ==", p[1:4]) // 3, 5, 7

	/* index 0 ~ 2 */
	fmt.Println("p[:3] == ", p[:3]) // 2, 3, 5

	/* index 4 ~ last */
	fmt.Println("p[4:] == ", p[4:]) // 11, 13
	fmt.Println("\n")
}

/*
 slice는 make 함수로 만들 수 있다. 생성된 slice는 0을 할당한 배열을 생성하고, 그것을 참조한다.
*/
func makeSlice() {
	a := make([]int, 5) // len(a) = 5
	printSlice("a", a)

	/* make 함수의 3번째 param을 통해 capacity 제한 가능 */
	b := make([]int, 0, 5) // len(b) = 0, cap(b) = 5
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	fmt.Println("\n")
}

/*
 빈 슬라이스
 슬라이스의 zero value는 nil이다.
 nil 슬라이스는 길이와 최대 크기가 0이다.
*/
func emptySlice() {
	var z []int
	fmt.Println("Z :", z, ", length:", len(z), ", capacity:", cap(z))
	if z == nil {
		fmt.Println("nil")
	}
}

func printSlice(s string, x []int) {
	// len = length, cap = capacity
	fmt.Printf("%s len = %d cap = %d %v \n", s, len(x), cap(x), x)
}

func main() {
	fmt.Println("[sliceBasic]")
	sliceBasicTest()

	fmt.Println("[slicing Slices]")
	slicingSlices()

	fmt.Println("[makeSlice]")
	makeSlice()

	fmt.Println("[emptySlice]")
	emptySlice()
}
