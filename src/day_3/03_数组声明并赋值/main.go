package main

import "fmt"
import "reflect"

// 1. 数组声明
// 数组的特点: (1). 数组固定长度的容器(一旦定好了数组长度。则不可变。即不可越界操作) (2). 数组内类型要求一致 (3). 数组是具有连续内存空间的容器。每个元素之间的相隔一个元素类型大小的字节(注：地址是以16进制数表示, 十六进制数是以字节为单位)。
// 数组的格式: var var_name [lenght]type
func arr_init() {
	// (1). 先声明, 后通过索引赋值
	var arr [3]int  // 数组的每个元素若声明但未初始化，则为类型的默认值
	fmt.Println(arr) // [ 0 0 0 ]
	fmt.Println(arr[0], arr[1], arr[2])
	// 通过索引赋值
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println(arr)

	// (2). 数组声明并初始化
	var names = [3]string{"kangknag", "wangming", "zhangsan"}
	fmt.Println(names) // [kangkang, wangming, zhangsan]
	var ages = [3]int{30, 40 ,50}
	fmt.Println(ages) // [30, 40, 50]

	// (3). 省略长度sh赋值
	var places = [...]string{"大连", "海口", "成都"}
	fmt.Println(places)

	// (4). 索引赋值
	var date = [...]string{0:"星期一", 2:"星期三"}
	fmt.Println(date)

	// 通过len()函数,可以计算容器数据的长度(可以查看容器中元素个数)
	fmt.Println(len("hello"))
	fmt.Println(len(places))
	fmt.Println(len("陈")) // 3, 一个汉字三字节
}

//  2. 数组的操作
func arr_operation() {
	// 操作1： 通过索引更改数组中的值
	// 声明并初始化数组
	var arr = [3]string{"kangkang", "wangming", "zhangsan"}
	// 通过索引更改数组中的值
	arr[0] = "李四"
	fmt.Println(arr)
	fmt.Println(reflect.TypeOf(arr))

	// 操作2: 切片操作: 数组[start索引:end索引]
	// 切片特点: (1).索引取值,左闭右开。(2).不像数组一样，有长度 
	var arr1 = [5]int{1, 2, 3, 4, 5}
	// 切片操作
	s1 := arr1[1:3]
	fmt.Println(s1, reflect.TypeOf(s1))
	s2 := arr1[1:]
	fmt.Println(s2, reflect.TypeOf(s2))
	s3 := arr1[:3] // 1,2,3 (因为切片的特性左闭右开。 所以索引只取到2的位置)
	fmt.Println(s3, reflect.TypeOf(s3))
	s4 := arr1[2:4]
	fmt.Println(s4, reflect.TypeOf(s4))

	// 操作3: 遍历数组操作
	var arr2 = [5]int{20, 21, 22, 23, 24}	
	// 1. 基础循环遍历
	for i := 0; i < len(arr2); i++ {
		fmt.Println(i, arr2[i])
	}
	// 2. 通过for range 进行循环遍历
	for index, value := range arr2 {
		fmt.Println(index, value)
	}
}

func main () {
	// arr_init()
	arr_operation()
}
