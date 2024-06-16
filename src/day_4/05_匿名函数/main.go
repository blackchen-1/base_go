package main

import "fmt"

/*
	匿名函数: 
		(1) 匿名函数：顾名思义, 就是没有名字的函数
		(2) 格式: func (形参变量, 形参变量) 返回值 {}    ---->  匿名函数本质就是一个函数体(拥有函数形参, 函数返回值)。
			1> 类似于c语言中的:
				c语言伪代码: 
					// 定义一个函数
					int add(int x, int y){
						return x + y; 
					}
				main中:
					//定义一个函数指针类型变量
					int (*func_p)(int, int)
					func_p = add
				注:	这个add使我们定义的函数，我们可以通过函数名的形式赋值给一个函数指针变量。也可以add(1, 2)这样进行执行

			2> 而GO语言中的匿名函数：
					匿名函数相当于add的操作。 唯一不同的就是,匿名函数是没有名字的。就是一个已经定义好的, 【函数参数(可以有,可无), 函数返回值(可有可无), 函数体】的一个整体。
					2.1> 它可以像c一样, 将其赋值给一个函数变量:
					var f = func(x int)(int){}

					2.2> 也可以像c一样,加一个括号(), 直接执行:
					(func(x int)(int){}) ()  -----> 匿名函数加一个括号, 将它们括起来, 是想表示它们(函数形参, 函数返回值, 函数体)是一个整体， 
*/

/*
	1. 匿名函数定义方法
*/
func anonymous_function() {
	// (1) 声明一个匿名函数, 并将其赋值给变量f
	var f = func(){
		fmt.Println("声明并将其赋值给新变量", "hello,world")
	}
	f() // 执行匿名函数

	// (2) 声明一个匿名函数, 并直接执行
	func() {
		fmt.Println("声明并执行", "hello, cwc")
	}()

	// (3) 声明一个有形参的匿名函数, 并直接执行
	func(x int, y int) {
		fmt.Println("有形参用例: ", x + y)
	}(10, 20)

	// (4) 声明一个有形参,有返回值的匿名函数, 并直接执行
	var ret_value int = (func (x int, y int) (int) {
		return x + y
	})(20, 30)
	fmt.Println(ret_value)
}


func main() {
	// 1. 匿名函数的使用
	anonymous_function()
}