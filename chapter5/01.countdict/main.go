package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	fmt.Println("开始数数")
	totalTime := 0
	for p := 0; p <= 5000; p++ {
		fmt.Println("正在统计第", p, "页")
		time.Sleep(10 * time.Second)
		r, _ := rand.Int(rand.Reader, big.NewInt(800))
		totalTime += r.Int()
	}
}
