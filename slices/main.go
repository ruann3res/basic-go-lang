package main

import "fmt"

func main() {
	var mySlice []int = []int{1, 2, 3, 4, 5}

	mySlice = append(mySlice, 6)

	fmt.Printf("Len=%d, Cap=%d, %v\n", len(mySlice), cap(mySlice), mySlice)
}
