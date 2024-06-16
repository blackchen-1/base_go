package main

import "fmt"

// 1. 引用类型（指针类型）相关讲解
// (1). 引用类型（指针类型）: 包括切片, map, channel都属于引用类型
// (2). var x *int （一级指针: 前面变量x, 指的是一个整型指针类型）
// (3). 区分一个概念:（1）. x是其值的地址，（2）.&x是x本身的地址。 若仅声明但未初始化x，则x的值地址为nil。
// (4). 地址是以16进制数表示, 十六进制数是以字节为单位。
// (5). 取址符( & ) ,取地址
// (6). 取值符( * ) ,取值
func point_print(){
	// 1. 声明并初始化x变量
	var x = 10
	fmt.Println("x的地址 :", &x)	// 通过&符号，取x的地址

	// 2. 声明一个指针类型变量, 存取x变量的地址
	var p1 *int // 指针变量 （一级指针） 
	p1 = &x
	fmt.Println("p1的地址", &p1) // p1的地址
	fmt.Println("p1的值(x的地址) :", p1) // x的地址
	fmt.Println("x地址的值(10) :", *p1) // 打印x的值

	// 3. 声明一个指针类型变量, 存取p的变量的地址
	var p2 **int // 指针变量 (二级指针)
	p2 = &p1 
	fmt.Println("p2的地址", &p2)	// p2的地址
	fmt.Println("p1的地址:", p2) // p2的值是p1的地址
	fmt.Println("p1地址的值(x的地址):", *p2) // p1的地址存的是x的地址
}

// 2. 值拷贝和引用拷贝的区别（DEMO）
func value_cp_and_point_cp() {
	// 值拷贝(值拷贝, 改y的值。看现象：x是否改变)
	x := 10
	y := x
	y = 20
	fmt.Println(x, y) // 10 , 20

	// 引用拷贝(通过引用拷贝, 改变y1指针的值(就是x1的地址)，指向的值。 看现象, x1是否改变)
	x1 := 100
	y1 := &x1 // y1 这里是一级指针，存的x1的地址
	*y1 = 200 
	fmt.Printf("x1的地址: %p \ny1的地址: %p \n", &x1, &y1)
	fmt.Println(x1, y1) // 200, y1是x1的地址 
}

// 指针类型练习
func practice() {
	var x int = 2
	y := &x
	*y++ // 虽然++的优先级是高于*。但是这里是会先*后++
	fmt.Println(y) // y = &x
	fmt.Println(x, *y)
}

func main() {
	// point_print()
	// value_cp_and_point_cp()
	practice()
}
