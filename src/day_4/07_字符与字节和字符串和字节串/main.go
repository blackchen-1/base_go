package main

import "fmt"
import "reflect"

/*
	一. 字符与字节
*/ 

func demo_1() {
	fmt.Println("--------------------------------------------")
	// 1. byte类型
	var y byte // 字节类型就是uint8类型
	y = 'a'
	fmt.Println(y, reflect.TypeOf(y)) // ascii码数字：97，类型就是uint8
	fmt.Printf("%d\n", y) // ascii码:97
	fmt.Printf("%c\n", y) // 字符: a
	fmt.Printf("%b\n", y) //97的二进制
	fmt.Printf("%x\n", y) // 97的十六进制

	fmt.Println("--------------------------------------------")
	// 2. uint8类型
	var b2 uint8
	b2 = 65
	fmt.Printf("%d\n", b2) // ascii码:65
	fmt.Printf("%c\n", b2) // 字符: A

	fmt.Println("--------------------------------------------")
	// 3. rune类型
	var z rune 
	z = '苑'
	fmt.Printf("字符 '苑' unicode的十进制: %d\n", z)
	fmt.Printf("字符 '苑' unicode的十六进制: %x\n", z)
	fmt.Printf("字符 '苑' unicode的二进制: %b\n", z)

	fmt.Println("--------------------------------------------")
	// 4. 字符串的简单使用
	var s = "苑abc"
	fmt.Println(s, len(s)) // 6个字节, 

	// 按照字节来取值，取得是一个字节的ascll码值
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) // 232 139 145 97 98 99 前三字节是存储中文苑的。 后三个字节分别存储a b c
	}

	// 按照一个字符来取
	for i, v := range s {
		fmt.Println(i, string(v))// 不加string，打印的是unicode值。加了string,打印的是utf8的值（直接输出中文，或者字符）
	}
}

/* 
	二. 字符串和字节串的转换
*/
func string_and_string() {
	// 定义一个字符串
	var msg = "abc苑"

	/* 
		一. 字符串转换成字节串
		注: 
		(1). 如果想看一个字符串中所有元素占多少个字节 -> 可以使用utf8的 []byte() 进行转换， 它输出的就是元素的utf8值
		(2). 如果想看一个字符串中有多少个元素，则可以使用[]rune()进行查看， 它输出的就是多少个元素的unicode的值
	*/
	fmt.Println([]byte(msg)) // 这个转换出来的值是utf8的值
	fmt.Println([]rune(msg)) // 这个转换出来的值是unicode的值

	/*
		二. 字节串转换成字符串
		(1). 下面两个info1和info2表达的都是一个意思， 第一个是将字符串转换成utf8的值，第二个就是只写写好的[]byte{}类型的utf8的值
	*/
	info1 := []byte(msg) 
	info2 := []byte{97, 98 , 99, 232, 139, 145}
	fmt.Println(info1)
	fmt.Println(info2)
	fmt.Println(string(info1))
	fmt.Println(string(info2))
}

func main() {
	// 一. 字符与字节
	// demo_1()

	// 二. 字符串和字节
	string_and_string()
}