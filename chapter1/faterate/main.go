package main

import "fmt"

func main() {
	var weight float64
	fmt.Print("体重(千克):")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高(米):")
	fmt.Scanln(&tall)

	var bmi float64 = weight / (tall * tall)

	var age int
	fmt.Print("年龄:")
	fmt.Scanln(&age)

	var sexWeight int
	fmt.Print("性别(男/女):")
	var sex string
	fmt.Scanln(&sex)

	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}

	var fatRate float64 = 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)

	fmt.Println("体脂率是", fatRate)

}
