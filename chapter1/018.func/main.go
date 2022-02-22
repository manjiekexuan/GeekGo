package main

import "fmt"

func guess(left, right uint) {
	guessed := (left + right) / 2
	var getFromInput string
	fmt.Println("我猜是：", guessed)
	fmt.Println("如果搞了，输入1，如果低了，输入0；对了输入9")
	fmt.Scanln(&getFromInput)

	switch getFromInput {
	case "1":
		guess(left, guessed-1)
	case "0":
		guess(guessed+1, left)
	case "9":
		fmt.Println("你心里想的数字是:", guessed)

	}

}

func main() {
	guess(1, 100)
}
