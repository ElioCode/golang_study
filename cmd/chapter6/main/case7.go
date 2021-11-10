package main

import "fmt"

func deferTest() {
	err := 1
	if err == 1 {
		fmt.Println("err != 1 : return")
		return
	}
	defer fmt.Println("defer")
	fmt.Println("err == 1 : return")
	return
}

func main() {
	deferTest()
}
