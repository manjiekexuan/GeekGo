package main

import "fmt"

func main() {
	var i int
	for i=1; i<=100;i++{
		fmt.Println("你好 golang")
	}

	j:=0
	for ;j<=5;j++ {
		fmt.Println("round2 hello golang")
	}

	k:=0
	for ;k<=5; {
		fmt.Println("round 3 ,hello golang")
	}
}
