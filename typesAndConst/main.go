package main

import (
	"fmt"
	"strconv"
)

// Types
// bool
// String
// int, int8 int16 int32 int64 (Positivos e negativos)
// unit uint8 uint16 uint32 uint64 uintptr (Apensa positivos)

// byte (Alias para o uint8)

// rune (Alias para int32) Representa um code point de um caracter Unicode

// float32 float64

// complex64 complex128

// string

func main() {
	const x = 64
	takeInt32(10) // 10
	takeInt32(x)  // 64

	s := strconv.FormatInt(int64(x), 10)
	intToString := string(x)

	fmt.Println(intToString) // @

	fmt.Println(s) // 64
}

func takeInt32(x int32) {
	fmt.Println(x)
}
