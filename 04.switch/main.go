package main

import "fmt"

func main() {
	var money int =20

	switch money{
	case 20:
		fmt.Println("点个外卖")
		fallthrough
	case 200:
		fmt.Println("下个馆子")
	case 2000:
		fmt.Println("容我想想")

	}
	fmt.Println("end")
}