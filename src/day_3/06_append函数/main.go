package main

import "fmt"



/*
	1. append 基本使用
*/
func append_base_using() {
	var s []int
	
	// (1). 在s的基础上，添加一个值
	s1 := append(s, 1) // 添加一个值: 1
	fmt.Println(s1)
	s2 := append(s1, 2, 3, 4) // 追加2,3,4
	fmt.Println(s2)

	// (2). 在s1的基础上,将s2这个切片中的元素追加到一起，赋值给s3切片
	// 如果要将s2中的元素追加到新的切片中, 需要切片后加...
	s3 := append(s1, s2...)
	fmt.Println(s3)

}

/*
	2. append 扩容机制
		(1). 容量不够, 两倍扩容, 新扩容的切片不与原来的切片共享一个底层数组。
			 若本身s容量不够, 给自身再使用append时候,虽然新扩容的本身容量。但该切片的地址不会变
		(2). 容量够, 与前一个切片共享一个数组。
*/
func append_expend_memory() {

	// 1. append扩容示例1
	var s = []int{1, 2, 3} // 长度是3，容量是3
	fmt.Println(s, "s len: ", len(s), "s cap: ", cap(s))
	s1 := append(s, 3) // 这时,原来的切片s的长度不够，新切片进行原来容量基础上两倍扩容。 与s不共享一个底层数组
	fmt.Println(s1, "s1 len: ", len(s1), "s1 cap: ", cap(s1))

	// 2. append扩容示例2
	var s2 = make([]int, 3, 5) // 长度为3, 容量为5
	fmt.Println(s2, "s2 len: ", len(s2), "s2 cap: ", cap(s2))
	
	s3 := s2[0 : 3] // 截取切片索引:0-2的元素赋值给s3。 此时两个变量共享一个数组
	s3[0] = 400 // 因为共享一个数组，s3改值, s2的第一个元素也改
	fmt.Println(s3, "s3 len: ", len(s3), "s3 cap: ", cap(s3))

	s4 := append(s3, 2000, 3000) // s4在s3的基础上追加两个元素。此时长度刚好的等于容量。 不进行扩容。 仍然共享一个数组
	s4[1] = 500
	s4[2] = 600
	fmt.Println(s4, "s4 len: ", len(s4), "s4 cap: ", cap(s4))

	s5 := append(s4, 600) // 此时， 容量不够了，进行扩容。 与之前的切片不共享一个底层数组。 长度: 6, 容量:10
	s5[0] = 99 // 验证与前面切片不共享一个数组。 因为共享一个数组这个切片改值，前面的值也会改
	fmt.Println(s5, "s5 len: ", len(s5), "s5 cap: ", cap(s5))

	// 3. append,切片本身扩容。其容量两倍扩容了，但地址不变
	s6 := []int{1, 2, 3}
	fmt.Printf("s6地址: %p\n", &s6)
	s6 = append(s6, 4)
	fmt.Printf("s6地址: %p\n", &s6)
}

/*
	3. append对切片的插入和删除
		(1). 头插
		(2). 任意位置插
		(3). 任意位置删除
		注: 将切片内容追加到另一个切片,需要加...
*/
func append_cdus_for_slice() {
	data := 999
	index := 3
	var s = []int{1, 2, 3, 4}
	fmt.Println(s)

	// 头插 (插入一个1)
	s1 := append([]int{1},  s...)
	fmt.Println(s1)

	// 任意位置插入
	s2 := append(s[:index], append([]int{data}, s[index:]...)...)
	fmt.Println(s2)

	// 任意位置删除
	s3 := append(s[:index], s[index+1:]...)
	fmt.Println(s3)

}

func main() {
	// // 1. append基本用法
	// append_base_using()

	// // 2. append扩容机制
	// append_expend_memory()

	// 3. append的增删改查
	append_cdus_for_slice()
}