package main

import "fmt"
import "strings"

// 练习: 获取账户名和密码的值
func getstr() (string) {
	s := "mysql... -u root -p 12345"
	return s
}

func main() {
	s := getstr()
	
	// 获取-u 和 -p 初始索引(就是-的索引)
	uindex := strings.Index(s, "-u")
	pindex := strings.Index(s, "-p")
	
	// 通过切片缺省获取值
	fmt.Println("账号: ", s[uindex + 2 : pindex - 1])
	fmt.Println("密码: ", s[pindex + 2 : ])
}