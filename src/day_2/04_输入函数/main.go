package main

import "fmt"
import "strings"

// io输入: Scan(必须将入参输入完，才结束输入)
// 注： 需要加上取址符(和C类似)
func input_Scan() {
	var name string
	var age int
	// 1. scan使用1
	fmt.Printf("请输入名字和年龄: ")
	fmt.Scan(&name, &age)
	fmt.Println(name, age)

}

// io输入: Scanln(不管入参是否输入完，换行即结束)
// 注: 需要加上取址符(与C类似)
func input_Scanln() {
	var name string 
	var age int 
	fmt.Printf("【名字】:")
	fmt.Scanln(&name)
	fmt.Println(name)

	fmt.Printf("【年龄】:")
	fmt.Scanln(&age)
	fmt.Println(age)
}

// io输入： Scanf(需要严格遵守,字符串中的格式输入,才能正确输入。 否则会报错)
// 注: 入参需要加上取地址符(与C类似)
func input_Scanf() {
	// var name string
	// var age int

	// fmt.Scanf("%s", &name)
	// fmt.Scanf("%d", &age)
	// fmt.Println(name, age)

	var x, y int

	fmt.Printf("按照这个格式输入两个值 [value_1 + value_2]- >")
	fmt.Scanf("%d + %d", &x, &y)
	fmt.Println(x / y)
}
// DEMO 
// 提示输入: 年-月-日。 
// 输出: xxxx年xx月xx日
func demo_Input() {
	var format_birth_data string

	fmt.Println("请输入您的生日, 按格式: 年-月-日")
	fmt.Scan(&format_birth_data)
	birthSlice := strings.Split(format_birth_data, "-")
	fmt.Println("birthSlice", birthSlice)
	fmt.Printf("您的生日是%s年-%s月-%s日\n", birthSlice[0], birthSlice[1], birthSlice[2])
}

func main() {
	// 1. IO输入Scan
	// input_Scan()
	// 2. IO输入Scanln
	// input_Scanln()
	// 3. IO输入Scanf
	input_Scanf()
	// DEMO1	
	// demo_Input()
}