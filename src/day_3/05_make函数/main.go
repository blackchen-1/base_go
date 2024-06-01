package main

import "fmt"

/*
	前言:
		指针和slice 这两个引用类型。当仅声明时，是不会开辟内存的。
		当然，如果直接声明并初始化赋值, 也是可以的。

	后言: 
	则new和make就应运而生
	(1). 前有new函数, 为仅声明但未初始化的指针，进行初始化
	(2). 后有make函数, 为引用类型的slice(切片), 进行初始化操作

	make注意点:
	例1: 
		var slice = make([]int, 5, 10)
	(1). 上面这是开辟了一个 len = 5, cap = 10 的整型切片
		也就是make开辟了开辟了10个整型的内存空间。 其中仅有5个空间是有默认值(int默认值: 0)。 
		剩余的5个空间，是已经开辟了,且可继续扩容的空间。
	
	例2: 
		var slice = make([]int, 5)
	(1). 上面这个切片slice, len = 5, cap在为给予固定值时,默认cap = len = 5
		也就是说当前这个切片开辟了5个可用的内存空间。 且有5个默认值元素。 所以 元素的总量不可以超过5。
*/

// 通过make进行初始化
func slice_init_by_make() {
	
	// // 1. 错误用例: 切片光声明但未初始化。便直接进行赋值操作
	// var slice []int // 仅声明但未初始化。 直接进行操作是不行的。会崩溃
	// slice[0] = 1
	// fmt.Println(slice)

	// 2. 除去上一个目录的直接声明并初始化赋值操作之外
	// 第二种初始化赋值切片的操作（make操作）
	var slice = make([]int, 5, 10) 
	slice[0] = 100 
	fmt.Println(slice) // [100 0 0 0 0 ]
}

func main(){
	slice_init_by_make()
}
