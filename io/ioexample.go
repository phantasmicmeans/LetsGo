package main

import (
	"fmt"
	"os"
)

func testmain() {
	fn := "C:/Users/SangMin/Desktop/golang/LetsGo/io/test.txt"
	withPanic(fn)
	withRecover(fn)
}

func withPanic(fn string) {
	fi, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	bytes := make([]byte, 1024)
	fi.Read(bytes)
	println("Content :", bytes)
	println("length :", len(bytes))
}

func withRecover(fn string) {
	defer func() { // execute after panic
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR")
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}
