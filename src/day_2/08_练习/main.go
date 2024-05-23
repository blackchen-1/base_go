package main

import "fmt"
import "strings"
import "strconv"
import "reflect"

/* 
	Pratice 1
	如何正确计算a+b的值
	var a int8 = 100
	var b int16 = 200
	fmt.Println(a + b)
*/ 
func Pratice_1() {
	var a int8 = 100
	var b int16 = 200
	// int8的范围是 -128 - 127 。所以不能用int8将b进行转换
	// 需要用int16将a转换成与b类型相同的才能相加
	fmt.Println(int16(a) + b)
}

/*
	pratice 2
	下面代码的执行结果是
	cmd := ""
	cmd += "l"
	cmd += "s"
	fmt.Println(cmd)
*/ 
func practice_2 () {
	// 字符串拼接
	cmd := ""
	cmd += "l"
	cmd += "s"
	fmt.Println(cmd)
}

/*
	3.	引导用户输入生日字符串，格式为“年-月-日” ，比如“1990-3-16”，
	然后按“您的生日是1990年-3月-16日”的格式化字符串输出到终端控制台。
*/
func practice_3() {
	var birthday string

	fmt.Printf("以此格式【年-月-日】输入 - >")
	fmt.Scanf("%s", &birthday)
	birthsilce := strings.Split(birthday, "-")
	fmt.Printf("您的生日是%s年-%s月-%s日\n", birthsilce[0], birthsilce[1], birthsilce[2])
}

/*
	4.  引导用户输入一个名字，判断该名字是否以小写a或者大写A开头，
		如果是打印true，否则打印false
*/ 
func practice_4(){
	var username string
	fmt.Printf("请输入你的名字 - >")
	fmt.Scan(&username)
	if (strings.HasPrefix(username, "A") || strings.HasPrefix(username, "a")) {
		fmt.Println("ture")
	} else {
		fmt.Println("false")
	}
}

/*
	5. 将整型数字100转换为字符串数字100
*/ 
func practice_5() {
	var int_value = 100
	str_value := strconv.Itoa(int_value)
	fmt.Println(str_value, reflect.TypeOf(str_value))
}

/*
	6. 将字符串"3.24"转为浮点数3.24
*/ 
func practice_6() {
	str_value := "3.24"
	float64_value, _:= strconv.ParseFloat(str_value, 64)
	fmt.Println(float64_value, reflect.TypeOf(float64_value))
}

/*
	7.  下面代码的执行结果是?，并分析why
		var a int8 = 100
		b := a
		a++
		fmt.Println("a:", a)// 101
		fmt.Println("b:", b)// 100
*/
func practice_7() {
	var a int8 = 100
		b := a
		a++
		fmt.Println("a:", a)// 101
		fmt.Println("b:", b)// 100
}

/*
	8. 某年是否为闰年可以依据“四年一闰，百年不闰，四百年闰”来进行判断。
	也就是说在能被 4 整除的年份当中，除了那些能被 100 整除但不能被 400 整除的年份外，其余的都是闰年。
	用户输入一个年份，判断是否为闰年
*/
func practice_8() {
	var year int
	fmt.Printf("请输入年份 - >")
	fmt.Scan(&year)
	if year % 4 == 0 {
		if (year % 100 == 0) && (year % 400 != 0)  {
			fmt.Printf("%d年 不是闰年\n", year)
		} else {
			fmt.Println("闰年")
		}
	} else {
		fmt.Println("不是闰年")
	}
}

/*
	9.	计算索引为10的斐波那契数列对应的值
	（斐波那契数列指的是一个数列 0, 1, 1, 2, 3, 5, 8, 13,特别指出：第0项是0，第1项是第一个1。从第三项开始，每一项都等于前两项之和）
*/
func practice_9(){
	var index_x = 0 //前两位
	var index_x_1 = 1 // 前一位
	var sum = 0
	for i := 2; i <= 10; i++ {
			sum = index_x + index_x_1
			index_x = index_x_1
			index_x_1 = sum
	}
	fmt.Println("索引为10的值 = ", sum)
}

/*
	10. 求1+2!+3!+4!+……+10!的和
	符号"!"表示阶乘运算。例如，n的阶乘（n!）表示从1到n的所有正整数相乘的结果。
*/
func practice_10() {
	var tmp int
	var sum = 0
	
	for i := 1; i <= 10; i++ {
		// 注意这里要将tmp，设置为1，否则每次迭代，都会保留上一次循环的值
		tmp = 1
		for j := 1; j <= i; j++ {
			tmp *= j
		} 
		sum += tmp
	}
	fmt.Println(sum)
}

/*
	11.	程序随机内置一个位于一定范围内的数字作为猜测的结果，由用户猜测此数字。
	用户每猜测一次，由系统提示猜测结果：太大了、太小了或者猜对了，直到用户猜对结果或者猜测次数用完导致失败。
	设定一个理想数字比如：66，让用户三次机会猜数字，如果比66大，则显示猜测的结果大了；如果比66小，则显示猜测的结果小了;
	只有等于66，显示猜测结果正确，退出循环。
	最多三次都没有猜测正确，退出循环，并显示‘都没猜对,继续努力’。
*/
func practice_11() {
	var target_value = 55 

	fmt.Println(`
规则如下:
	1. 在范围 1 - 100内猜数字
	2. 三次机会， 其中每猜错一次会提示值是否大于还是小于目标值 
	`)

	var input_value int
	var count = 3
	for true {
		fmt.Printf("输入对应值 - > ")
		fmt.Scan(&input_value)
		if input_value > 100 && input_value < 0 {
			fmt.Println("非法输入,请重试...")
		} else if input_value > target_value {
			count--
			fmt.Printf("输入值大于目标值,还有%d次机会\n", count)
		} else if input_value < target_value {
			count--
			fmt.Printf("输入值小于目标值,还有%d次机会\n", count)
		} else {
			fmt.Printf("猜对了，你太牛逼了！！！\n")
			break
		}

		if (count == 0) {
			fmt.Println("机会用尽， 下次加油！！！\n")
			break
		}
	}
}

func main() {
	// practice_1()
	// practice_2()
	// practice_3()
	// practice_4()
	// pratice_5()
	// practice_6()
	// practice_7()
	// practice_8()
	// practice_9()
	// practice_10()
	practice_11()
}
