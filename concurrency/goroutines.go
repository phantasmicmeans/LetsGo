package main

import (
	"fmt"
	"time"
)

/*
A goroutine is a lightweight thread managed by the Go runtime.
	go f (x, y, z)
starts a new goroutine running
	f (x, y, z)
The evaluation of f (x, y, z) happens in the current goroutine
and the executions of "f" happens in the new goroutine.

Goroutines run in the same address spaces,
so access to shared memory must be synchronized.

The sync package provides useful primitives,
although you won't need them much in Go as there are other primitives.
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}