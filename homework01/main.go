package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var inputStr string
	// for 循环初始值
	var count int = 3

	var sexWeight int
	//循环输入次数
	fmt.Println("姓名;性别(男,女);身高(米);体重(公斤);年龄 ------> 例如:小治;男;1.81;86.5;30")

	for i:=1;i<=count;i++ {
		fmt.Printf("请输入第%v个人的信息",i)
		fmt.Println()
		fmt.Scanln(&inputStr)
		input := strings.Split(inputStr,";")
		//判断参数个数是否正确
		if len(input) < 5 || len(input) > 5 {
			fmt.Println("输入的参数个数不正确请重新输入")
			count++
		}
		name := input[0]
		sex  := input[1]
		if sex == "男"{
			sexWeight = 1
		}else if sex == "女"{
			sexWeight = 0
		}
		height, err := strconv.ParseFloat(input[2], 64)
		if err != nil {
			fmt.Println("输入的身高不正确请重新输入")
			count++
		}
		weight, err := strconv.ParseFloat(input[3], 64)
		if err != nil {
			fmt.Println("输入的体重不正确请重新输入")
			count++
		}
		age,err:= strconv.ParseInt(input[4],10,0)
		if err != nil {
			fmt.Println("输入的年龄不正确请重新输入")
			count++
		}
		BMI,FatRate,device:= getPersonReport(weight,height,age,sexWeight)
		fmt.Println(name,BMI,FatRate,device)

	}
	fmt.Println("已经输入结束,报表如下")
}


func getPersonReport(weight float64,height float64,age int64, sexWeight int )(float64,float64,string){
	BMI:=getBMI(weight,height)
	fatRate := getFatRate(BMI,age,sexWeight)
	return BMI,fatRate,getDevice(fatRate,age,sexWeight)

}

func getDevice(fatRate float64,age int64,sexWeight int )string{
	//男性体脂
	if sexWeight == 1{

	}
}

func getBMI(weight float64,height float64) float64{
	//bmi := weight / (height * height)
	return weight / (height * height)
}

func getFatRate(BMI float64,age int64,sexWeight int) float64{
	//fatRate := 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)
	return 1.2*BMI + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)
}