package main

import "io/ioutil"

func main() {
	fn := "C:/Users/SangMin/Desktop/golang/LetsGo/io"
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("RECOVER")
		}
	}

	bytes, err := ioutil.ReadFile(fn + "/test.txt")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fn + "/test2.txt", bytes, 0)
	if err != nil {
		panic(err)
	}
}
