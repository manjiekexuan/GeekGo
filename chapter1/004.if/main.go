package main

import "fmt"

func main() {
	main1()
	/*	var fruit string = "6 apples"
		var watermallan bool = true // true or false
		if watermallan {
			fruit = "1 apple"
		} else {
			fruit = "7 apples"
		}
		fmt.Println("buy: ", fruit)*/
}

func main1() {
	var fruit string = "6 apples"
	var watermallan bool = false
	if watermallan {
		fruit = " 1 apple"
	} else {
		fruit = "7 apples"
	}
	fmt.Println("buy:", fruit)
}
