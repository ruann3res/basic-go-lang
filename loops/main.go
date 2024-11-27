package main

import "fmt"

// for i :=0 (Init statement); i < 10 (Condition); i++ (Post statement) {
// init statement: Executado antes da primeira iteração do loop
// condition: É avaliado antes de cada iteração do loop
// post statement: Executado no final de cada iteração do loop

func main() {
	arr := [3]int{1, 2, 3}
	for index, element := range arr {
		fmt.Println(&index, &element)
	}
}
