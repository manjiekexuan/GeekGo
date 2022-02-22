package main

import "fmt"

func main() {
	//names := []string{"小强", "周毅", "旭东"}
	//fr := []float64{28, 18, 18}
	//
	//names = append(names, "张松")
	//fr = append(fr, 19)
	//
	//for i, name := range names {
	//	if name == "周毅" {
	//		fmt.Printf("%s 的体脂率是 %f\n", name, fr[i])
	//	}
	//}

	var m1 map[string]int
	m2 := map[string]int{}
	m3 := map[string]int{"王强": 60, "李静": 83, "苗文": 91}
	fmt.Println(m1, m2, m3)
}
