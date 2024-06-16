package main

import "fmt"
import "reflect"

/*
	字符与字节
*/ 

func demo() {
	// byte类型
	var y byte // 字节类型就是uint8类型
	y = 'a'
	fmt.Println(y, reflect.TypeOf(y)) // ascii码数字：97，类型就是uint8
	fmt.Printf("%d\n", y) // ascii码:97
	fmt.Printf("%c\n", y) // 字符: a
	fmt.Printf("%b\n", y) //97的二进制
	fmt.Printf("%x\n", y) // 97的十六进制

	// uint8类型
	var b2 uint8
	b2 = 65
	fmt.Printf("%d\n", b2) // ascii码:65
	fmt.Printf("%c\n", b2) // 字符: A
}

func main() {
	demo()
}