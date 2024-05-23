package main

import "fmt"

// switch - case （go的switch,不同于c。不需要break跳出。）
func DEMO_switch_case() {
	var week int
	fmt.Print("请输入星期值 -> ")
	fmt.Scan(&week)
	switch (week) { // go中这里可以不加括号，C中需要添加
	case 1:
		fmt.Println(" 星期一 ")
	case 2:
		fmt.Println(" 星期二 ")
	case 3:
		fmt.Println(" 星期三 ")
	case 4: 
		fmt.Println(" 星期四 ")
	case 5:
		fmt.Println(" 星期五 ")
	case 6:
		fmt.Println(" 星期六 ")
	case 7:
		fmt.Println(" 星期天 ")
	default:
		fmt.Println(" 非法输入 ")	
	}
}

func main() {
	DEMO_switch_case()
}