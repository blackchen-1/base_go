package main

import "fmt"

// 1. 函数的常规使用(通过传参, 实现功能)
func add(x int, y int) {
	fmt.Println(x + y)
}

/* 
	2. 函数通过定义, 不定长的形参, 实现功能
	注: 定义不定长的形参时, 实际是定义了一个切片类型的形参。将多个实际参数存入切片中,进而索引取值,实现函数功能
	形参定义格式: 变量名 + ... + 变量类型
*/
func add_more_num(s ... int){ // 不定长形参
	var ret int = 0
	for _, v := range s {
		ret += v
	}
	fmt.Println(ret)
}

/*
	3. 函数返回值(前面定义的函数基本都是无返回值的,所以没有限制返回值形参)
		注: 
		(1). 如果函数需要返回值，则需要规定函数返回值类型
		(2). 如果返回一个值可以不用加括号，如果返回多个值需要加括号
*/
// func login(uname string, pwd string) bool { // 这是返回一个参数可以不加括号
func login(uname string, pwd string) (bool, string) {
	if uname == "张三" && pwd == "12345" {
		return true, uname
	} else {
		return false, ""
	}
}

/*
	4. 函数返回值命名
	注：
	(1). (若对函数返回值命名了, 则该函数结束, 只返回命名函数的值。 若未对该定义返回变量操作,则返回该变量类型默认值)
*/
func demo(v1 int, v2 string) (x int, y string) {
	x = v1
	y = v2
	return // 等于 return x, y
}

func main() {
	// 1. 函数通过传参,实现功能
	add(1, 2)

	// 2. 函数通过传入多入实际参数,实现功能
	add_more_num(1, 2, 3, 4)

	// 3. 函数返回值(返回一个或多个返回值)
	isok, uname := login("张三", "12345")
	fmt.Println(isok, uname)

	// 4. 函数返回值命名
	x, y := demo(3, "11111")
	fmt.Println(x, y)
}
