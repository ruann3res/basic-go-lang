package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var myCondition int = 1

	switchCase(myCondition)
	switchVariable()

	fmt.Println(isWeekend(time.Now()))

	compareType(10)      // 10
	compareType("Hello") // Hello
}

func switchCase(condition int) {

	switch condition {
	case 1:
		fmt.Println("Case 1")
		fallthrough
	case 2:
		fmt.Println("Case 2")
	default:
		fmt.Println("Default")
	}
}

func switchVariable() {
	switch x := math.Sqrt(4); x {
	case 2:
		fmt.Println("Resultado é 2")
	default:
		fmt.Println("Algo esta errado")
	}
}

func isWeekend(x time.Time) bool {
	switch x.Weekday() {
	case time.Sunday, time.Saturday:
		return true
	default:
		return false
	}
}

func compareType(x any) {
	switch t := x.(type) {
	case int:
		takeInt(t)
	case string:
		takeString(t)
	default:
		fmt.Println("Default")
	}
}

func takeString(s string) {
	fmt.Println(s)
}

func takeInt(x int) {
	fmt.Println(x)
}
