package main

import "fmt"
import "time"

func foo() {
	fmt.Println("foo功能开始")
	time.Sleep(time.Second * 2)
	fmt.Println("foo功能结束")
}

func bar() {
	// 时间戳
	fmt.Println("bar功能开始")
	time.Sleep(time.Second * 3)
	fmt.Println("bar功能结束")
}

/* 
	一. 高阶函数以函数作为参数
	DEMO:
 	定义一个程序运行时间的函数, 以函数作为参数，传参。测试函数运行时间
*/
func timer(f func()) {
	start := time.Now().Unix()
	f()
	end := time.Now().Unix()
	fmt.Println("cost timer:", end-start)
}

/*
	二. 高阶函数以函数为返回值

*/
func func_return_type_func() func() int {
	// 声明一个匿名函数
	var inter = func() int {
		fmt.Println("hello,world")
		return 0
	}
	return inter
}


func main() {
	// // 一. 高阶函数以函数作为参数
	// timer(foo)
	// timer(bar)

	// 二. 以函数作为返回值
	var f = func_return_type_func() 
	fmt.Println(f())
}