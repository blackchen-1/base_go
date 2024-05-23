package main

import "fmt"

// 循环 - for 常用方式1
func  loop_for_normal_using_1() {

	num := 100  	// 初始语句
	for num > 0 {  	// 条件表达式
		fmt.Println(num)
		num-- 		// 步进语句
	}
}

// 循环 - for 常用方式2
func loop_for_normal_using_2() {
	for count := 0; count < 10; count++ {
		fmt.Println("hello,wolrd")
	}
}

// 循环退出 - for-break , for-continue
// 注: continue跳过当次循环， break 退出单个循环
func loop_quit() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j ++ {
			if (j == 5) {
				fmt.Printf("j = %d时, 跳过这次循环\n", j)
				continue
			}
			if (j == 7) {
				fmt.Printf("当 j = %d时, 退出j循环\n", j)
				break
			}
			fmt.Println(j)
		}

		if (i > 6) {
			fmt.Println(`
				i > 6 时
				执行break: 
					退出i循环
			`)
			break
		}
	}
}

// TEST循环的问题 1- for循环局部变量之间作用域问题（（不建议这样定义变量））
func local_var_in_same_func() {
	for i := 0; i < 3; i++ { // 内部能够拿到外部的循环的值。 但外部无法拿到内部循环的值。所以外部循环的同名变量不收影响
		fmt.Println("相同的for循环作用域下: 循环外同名i", i)
		for i := 0; i < 3; i++ {
			fmt.Println("不同的for循环作用域下: 循环内同名i", i)
		}
	}

	for i := 0; i < 3; i++ { // 虽然相同函数作用域, 但for循环初始语句定义的变量具有局部性,所以两个同名变量之间不受影响
		fmt.Println("相同函数作用域下,不同for循环循环作用域下: 循环内同名i", i)
	}
}

// DEMO1 : 基于循环打印: 1+2+3+4+...+100
func DEMO_1() {
	sum := 0
	for num := 0; num <= 100; num++ {
		sum += num
		fmt.Println(sum)
	}
	fmt.Println("最终和: ", sum) 
}

func main() {
	// 循环for, 常用方式 1
	// loop_for_normal_using_1()
	// 循环for, 常用方式 2
	// loop_for_normal_using_2()
	// 退出循环
	loop_quit()
	//  DEMO 1
	// DEMO_1()
	// 	TEST 1
	// local_var_in_same_func()
}
