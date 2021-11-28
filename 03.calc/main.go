package main

import "fmt"

func main() {
	var a,b = 3,10

	fmt.Println("a+b=",a+b)
	fmt.Println("a*b=",a*b)
	fmt.Println("a/b=",a/b)
	fmt.Println("a%b=",a%b)


	var A,B,C bool = true,false,false

	fmt.Println(A&&(B=C))
}