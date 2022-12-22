package main

import (
	"fmt"
	"time"
)

func main() {
	// declaracion 1
	var a int
	a = 10
	fmt.Println(a)

	// declaracion 2
	b := 10
	fmt.Println(b)

	// tipos principales pero existen mas
	var a1 int
	var a2 string
	var a3 bool
	var a4 time.Time // fecha con hora, ej: 2022-10-01 11:30

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
}
