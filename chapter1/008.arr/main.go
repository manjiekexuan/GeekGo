package main

import "fmt"

func main() {
	/*	xqInfo := []string{"小强", "男", "在职"}
		xl := []string{"小李", "男", "在职"}
		xsInfo := []string{"小苏", "女", "在职"}*/
	var xueye string
	fmt.Println("请输入薛也最讨厌的人:")
	fmt.Scanln(&xueye)

	if xueye == "陈莹" {
		fmt.Println("已去世多时")
	} else if xueye == "大亩0" {
		fmt.Println("和严鹏飞做姐妹")
	}

}
