package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fn := "C:/Users/SangMin/Desktop/golang/LetsGo/io"
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("RECOVER")
		}
	}()

	f, err := os.Open(fn + "/test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fo, err2 := os.Create(fn + "/test2.txt")
	if err2 != nil {
		panic(err2)
	}
	defer fo.Close()

	buff := make([]byte, 1024)
	for {
		cnt, err := f.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if cnt == 0 { // 끝이면 루프 종료
			break
		}

		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}
}
