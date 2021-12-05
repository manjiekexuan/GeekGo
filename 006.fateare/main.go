package main

import "fmt"

func main() {
	var weight float64
	fmt.Print("体重(千克)")
	fmt.Scanln(&weight)


	var tall float64
	fmt.Print("身高(米)")
	fmt.Scanln(&tall)

	var sexWeight int
	fmt.Print("性别")
	fmt.Scanln(&sexWeight)

	var bmi float64 = weight / (tall * tall)
	var age int = 30
	var fatRate float64 = 1.2 * bmi + 0.23 * float64(age) - 55.4 -10.8 * weight

}
