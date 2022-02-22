package main

import (
	"fmt"
)

func main() {
	var left, right int
	var op string
	fmt.Scanln(&left)
	fmt.Scanln(&right)
	fmt.Scanln(&op)

	newC := NewCalculator{}
	newC.left = 100
	newC.right = 200
	fmt.Println(newC.Add())
}
