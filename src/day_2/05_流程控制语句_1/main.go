package main

import "fmt"

// 单分支语句 if 
func branch_single() {
	name := "kangkang"

	if (name == "wangming") {
		fmt.Println("name_1 正确 !!!")
	}

	if (name == "kangkang") {
		fmt.Println("name_2 正确 !!!")
	}
	// 注: go支持 if后面的表达式不加括号(和python一样)。 C中是需要加括号 
}

// 双分支语句 if - else
func branch_double() {
	var name string
	fmt.Printf("【姓名】->")
	fmt.Scanln(&name)

	if name == "kangkang" {
		fmt.Println(" name 正确 !!!")
	} else {
		fmt.Println(" name 错误 ...")
	}
}

// 多分支语句 if - else if - else 
func branch_many() {
	local_name := "kangkang"
	remote_name_1 := "kangkang"
	remote_name_2 := "wangming"
	if local_name == remote_name_1 {
		fmt.Println("远程有该用户", local_name)
	} else if local_name == remote_name_2 {
		fmt.Println("远程有该用户", local_name)
	} else {
		fmt.Println("远程没有该用户", local_name)
	}
}
// DEMO 1
// 终端输入用户名和密码 
// 判断用户名及密码是否正确
func DEMO_if_else(username string, password string) {
	local_username := "kangkang"
	local_password := "12345"

	if ((local_username == username) && (local_password == password)) {
		fmt.Println("登录成功 ！！！")
	} else {
		fmt.Println("用户名或密码错误, 请检查后再输入...")
	}

}

// DEMO 2
func DEMO_if_elseif_else(score int) {
	if score < 0 || score > 100{
		fmt.Println("成绩不合理")
	} else if score > 90 {
		fmt.Println("成绩优秀")
	} else if score > 60 {
		fmt.Println("成绩合格")
	} else {
		fmt.Println("成绩不合格")
	}
}

func main() {
	// 1. 单分支语句 if
	// branch_single()
	// 2. 双分支语句 if - else
	// branch_double()
	// 3. 多分支语句 if - else if - else
	branch_many()

	// 4. DEMO 1
	// var username, password string
	// fmt.Printf("请输入用户名->")
	// fmt.Scan(&username)
	// fmt.Printf("请输入密码->")
	// fmt.Scan(&password)
	// DEMO_if_else(username, password)

	// 5. DEMO 2
	// var  score int
	// fmt.Printf("请输入你的成绩 - > ")
	// fmt.Scan(&score)
	// DEMO_if_elseif_else(score)

}
