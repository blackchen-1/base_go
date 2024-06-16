package main

import "fmt"

/*
	函数的值传递和引用传递
*/

/*
	1. 函数值传递用例
		函数的值传递,并在函数中作操作,是不会改变传递变量本身的
*/
func value_cp(s int) {
	s += 1
}

/*
	2. 函数的引用传递 示例1(传递变量地址)
		函数的引用传递,就是对指针的操作。通过取地址，对地址存的值进行修改。是可以改变变量本身的值的
*/
func point_cp(s *int) {
	*s += 1
}

/*
	3. 函数的引用传递 示例2(传递切片)
		(1) 首先切片本身组成 ： 【底层数组的首地址, 切片长度, 切片容量】
		(2) 在切片传递时,实际把切片的底层数组的首地址,以及切片长度,容量。 
		相当于拷贝给了函数的形参slice。但这个形参的切片和实参的切片本质不是一个地址。不过是存的元素地址是相同的。可以共同操作一个底层数组
*/ 
func slice_cp(slice []int) {
	slice[0] = 200 // 会改变实参切片值
	fmt.Println("形参slice :", slice)
	slice1 := append(slice, 999) // 原来实参容量不足,此时append两倍扩展了新的空间，slice和slice1不共享一个底层数组。所以改值不会改变实参s的切片值
	slice1[0] = 300
	fmt.Println("slice1 :", slice1)
}

func main() {
	// 1. 函数的值传递(传递值变量本身)
	var v1 int = 20
	value_cp(v1)
	fmt.Println("v1: ", v1)

	// 2. 函数的引用传递(传递变量的地址)
	var v2 int = 20
	point_cp(&v2)	
	fmt.Println("v2: ", v2)

	// 3. 函数的引用传递(传递切片)
	var s = []int{1, 2, 3}
	slice_cp(s)
	fmt.Println("实参切片: ", s)
}
