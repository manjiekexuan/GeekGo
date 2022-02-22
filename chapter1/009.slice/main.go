package main

import "fmt"

func main() {
	a := []int{}

	fmt.Println(a)
	fmt.Println("追加元素到a中，a是切片")

	a = append(a, 333)

	fmt.Println(a)

	//b := [0]int{}

	fmt.Println("删除")

	a = []int{111, 222, 333, 444, 555}

	b := append(a[:2], a[3:]...)

	fmt.Println(b)
}
