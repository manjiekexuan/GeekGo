package main

import "fmt"

func main() {
	bmis := []float64{1.2, 3.222, 4.342344}

	avgBmi := calculateAvg(bmis)

	fmt.Println(avgBmi)
}

func calculateAvg(bmis []float64) (avgBmi float64) {
	//var total float64 = 0
	for _, item := range bmis {
		fmt.Println(item)
		//total += item
	}
	//return total / float64(len(bmis))
	return 13.2
}
