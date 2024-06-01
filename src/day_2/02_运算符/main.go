package main

import "fmt"
// c语言的运算符总结: 括号，单算移关与，异或逻条赋，逗号
// go语言的运算符优先级总结:
func opa_precedence() {
	fmt.Println(`
运算符的优先级见本目录下png文件:
	"day_2/02_运算符/元运算符优先级.png"
		`)
}

// 1. 科学运算符 + - * / %
func opa_science() {
	fmt.Println(`1. 科学运算符`)
	var x = 10
 	var y = 20
 	fmt.Println(x + y)
 	fmt.Println(x - y)
 	fmt.Println(x + y)
 	fmt.Println(x + y)
 	fmt.Println(x + y)
}

// 2. 比较运算符 > < == >= <= != 
func opa_compare() {
	fmt.Println(`2. 比较运算符`)
	var x = 10
	var y = 20
	fmt.Println(x < y)
	fmt.Println(x > y)
	fmt.Println(x == y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x != y)
}

// 3. 逻辑运算符 &&(与)、||(或)、!(非/取反)
func opa_logic() {
	fmt.Println(`3. 逻辑运算符`)
	var x = true
	var y = false
	fmt.Println(x && y) // f
	fmt.Println(x || y) // t
	fmt.Println(!x) // f
	fmt.Println(!y) // t
}

// 4. 赋值运算符
func opa_AssignMent(){
	fmt.Println(`4. 赋值运算符`)
	var x int 
	x = 1  // = 
	x += 2 // x = x + 2
	x -= 2 // x = x - 2
	x *= 2 // x = x * 2
	x /= 2 // x = x / 2 
	x %= 2 // x = x % 2
	x++ 	// 自增运算符 x = x + 1 
	x--		// 自减运算符 x = x - 1
	// 注: 不可x = x++ 或 x = x--
 	fmt.Println(x)
}

// DEMO: 判断奇数和偶数
func is_Even_Or_Odd(x int) (string) {
	fmt.Println(`5. DEMO:`)
	if (x % 2 == 0) {
		return "偶数"
	} else {
		return "奇数"
	}
}

// DEMO: 字符串校验(使用逻辑运算符)
func vali_Uname_Pwd(username string, password string) (string){
	remote_username := "kangkang"
	remote_password := "12345"
	if ((remote_username == username) && (remote_password == password)) {
		return "校验成功"
	} else {
		return "校验失败"
	}
}

func main() {
	/* knowledge */
	opa_precedence() // 运算符优先级
	opa_science() // 科学运算符知识点
	opa_compare() // 比较运算符知识点
	opa_logic() // 逻辑运算符知识点
	opa_AssignMent() // 赋值运算符知识点
	/* DEMO */ 
	fmt.Println(is_Even_Or_Odd(20)) // 判断是偶数还是奇数
	fmt.Println(vali_Uname_Pwd("wangming", "12345")) // 字符串校验
	fmt.Println(vali_Uname_Pwd("kangkang", "12345"))
}