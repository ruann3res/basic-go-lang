package main

import "fmt"

func main() {
	salarios := map[string]float64{
		"Maria": 11325.45,
		"João":  11325.45,
	}
	fmt.Println(salarios["Maria"]) // Maria:11325.45]

	salarios["Pedro"] = 1200.0
	fmt.Println(salarios) // map[João:11325.45 Maria:11325.45 Pedro:1200]
}
