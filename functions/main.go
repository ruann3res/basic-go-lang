package main

import (
	"fmt"
	// "module1/names/internal/mine" Nao e possivel importar pois o pacote e privado
)

func main() {
	fmt.Println("Go project")
	simpleResult := hiOrderSum(1)(2)
	sumFunction := hiOrderSum(4)
	sumValue := sumFunction(1)
	fmt.Println(simpleResult)
	fmt.Println(sumValue)
	f := variaticSum(10, 20, 30)
	fmt.Println(f)
}

func sayHi() {
	fmt.Println("Hi")
}

func sum(a int, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func subtract(a, b int) int {
	return a - b
}

func division(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return
}

func swapValue(a, b int) (int, int) {
	return b, a
}

func hiOrderSum(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func variaticSum(nums ...int) int {
	var out int
	for _, num := range nums {
		out += num
	}
	return out
}
