package main

import (
	"fmt"
)

func main() {
	arr := [13]int{5: 400}
	fmt.Println(arr) // [0 0 0 0 0 400 0 0 0 0 0 0 0]

	arr2 := [13]string{5: "400", 7: "Ola Mundo"}
	fmt.Println(arr2) // [  400     Ola Mundo]

	const arrSize = 15
	arr3 := [arrSize]int{5: 400}
	fmt.Println(arr3) // [0 0 0 0 0 400 0 0 0 0 0 0 0 0 0]

	for i, v := range arr3 {
		fmt.Println(i, v)
	}
}
