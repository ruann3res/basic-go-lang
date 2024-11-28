package main

import "fmt"

func main() {
	x := doDefer()
	fmt.Println(x)
}

func doDefer() int {
	defer fmt.Println("world")

	fmt.Println("Hello")

	return 10
}
