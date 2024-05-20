package main

import "fmt"

// 1. fmt.Println(向终端输出、自带换行、 可以直接不用占位符、直接变量输出)
func outPut_Println() {
	var name, age = "world", 20
	fmt.Println("hello", name, age)
}

// 2. 标准输出 fmt.Printf(向终端输出、使用占位符)
func outPut_Printf() {
	var name, age, isMarried = "kangkang", 20, false
	fmt.Printf("姓名: %s , 年龄： %d, 婚否: %t \n", name, age, isMarried)
}

// 3. Sprintf (拼接各类型值, 返回字符串)
func outPut_Sprintf() {
	var name, age, isMarried = "kangkang", 20, false
	var s string = fmt.Sprintf("姓名: %s , 年龄: %d , 婚否: %t", name, age, isMarried)
	fmt.Println(s)
	// 注： 这个Sprintf是不向终端输出的。而是返回字符串
}

func main() {
	// 向终端输出
	outPut_Println()
	outPut_Printf()
	// 返回字符串
	outPut_Sprintf()
}
