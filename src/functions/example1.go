package main

import (
	"fmt"
)

func main() {
	fmt.Println(suma(10, 10))
	fmt.Println(sumaWithTwoResponse(10, 10))
}

func suma(x int, y int) int {
	r := x + y
	return r
}

func sumaWithTwoResponse(x int, y int) (response int, firstNumber int) {
	response = x + y
	firstNumber = x
	return
}
