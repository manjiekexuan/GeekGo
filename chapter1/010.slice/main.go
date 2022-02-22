package main

import "fmt"

func main() {
	var a string = "你好"

	fmt.Println(a)

	aBytes := []rune(a)

	fmt.Println(aBytes)

	fmt.Println("修改切片内的内容")

	aBytes[0] = 'H'

	a = string(aBytes)

	fmt.Println(a)

}
