package main
import "fmt"

func var_use_1() {
	// 1. 声明变量
	var age int
	fmt.Println(age)
	// 2. 变量赋值
	age = 22
	fmt.Println(age)
}

func var_use_2() {
	// 声明并赋值一行字符串实现
	// 1. 分两步实现声明并赋值（显式声明）
	// 声明字符串变量
	var name_1 string
	// 赋值字符串变量
	name_1 = "kangkang"
	fmt.Println(name_1)

	// 2. 一步实现声明并赋值(隐式声明)
	var name_2 = "kangkang"
	fmt.Println(name_2)

	// 3. 一步实现声明并赋值(隐式声明以及:= 进行声明)
	// 注: 当全局声明时，不能使用:=进行声明并赋值
	name_3 := "kangkang" 
	fmt.Println(name_3)

	// 4. 声明并赋值多个变量
	name_4, age, isMarried := "kangkang", 23, false
	fmt.Println(name_4, age, isMarried)
}

func var_use_3() {
	// 一次声明多个变量
	// 法1:
	// var x, y int
	// var x, y string
	// var x, y bool
	// fmt.Println(x, y)

	// 法2：
	var (
		a int  // 0
		b string  // ""
		c bool // false
	)
	fmt.Println(a, b, c)
}

func main() {
	// 一. 变量声明与赋值
	var_use_1()
	// 二. 声明并赋值（其中涉及知识点: 显隐式声明的区分）
	var_use_2() 
	// 三. 一次声明多个变量(其中涉及知识点: 变量的默认值)
	var_use_3()
}
