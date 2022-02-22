package main

import "fmt"

func main() {
	var age int = 35
	fmt.Println(age)

	// 初始化数组
	var ages [5]int = [5]int{35, 32, 33, 37, 59}
	fmt.Println(ages)

	var ages2 = [5]int{35, 32, 33, 37, 59}
	fmt.Println(ages2)

	ages3 := [5]int{35, 32, 33, 37, 59}
	fmt.Println(ages3)

	age4 := [3]int{}
	age4[0] = 1
	age4[1] = 2
	age4[2] = 3
	print(age4)
}
