package main

import "fmt"
import "strconv"
import "reflect"

// 整型与整型之间的转换
func int_to_int() {
	var a int8 = 10
	var b int16 = 20
	// fmt.Println(a + b) // 这样是不行的,因为类型不同
	fmt.Println(a + int8(b)) // 通过int8()函数,将int16类型的b转成int8
}

// 字符串转整型
func string_to_int() {
	// 法1: strconv.Atoi
	var agestr = "100"
	var ageint, _ = strconv.Atoi(agestr)// 返回值是两个,转换后的int值和错误log:error
	fmt.Println(reflect.TypeOf(ageint))
} 

// 整型转字符串
func int_to_string() {
	price := 100
	priceStr := strconv.Itoa(price)
	fmt.Println(price, reflect.TypeOf(priceStr))
}

// 以特殊的转换函数
func parse_num_api() {
	// 1. 字符串转换为整型(第二个参数是确定字符串中的值是几进制, 第三个参数是限制了值的bit的大小)
	// 假设是第三个参数是 8, 则转换返回值范围则是有符号的int8,就是[-128, 127], 如果值超出范围则会报错误
	// 一般第三个参数都放64
	// num_int64, error := strconv.ParseInt("011", 2, 64) // 二进制输入, 返回int64类型值
	// num_int64, error := strconv.ParseInt("4E20", 16, 64) // 十六进制输入, 返回int64类型值
	num_int64, error := strconv.ParseInt("200", 10, 64) // 十进制输入, 返回int64类型值
	fmt.Println(num_int64, "\terr:", error)
	fmt.Println(reflect.TypeOf(num_int64))
	
	// 2. 字符串转换为浮点型(浮点型有float32(单精度) 和 float64(双精度) 两个类型)
	// 第二个参数与精度有关, 64表示输入类型为float64位，32表示输入类型为float32位。
	// 但无论输入是float32还是float64,都会转化为64位进行返回
	num_float_64, error := strconv.ParseFloat("3.14159262222", 64)
	fmt.Println(num_float_64, "\terr:", error)
	fmt.Println(reflect.TypeOf(num_float_64))

	// 3. 字符串转换为bool类型(将某些值转换为bool值。 仅限于"0, t, T, true, false...等等。如果不相关的输入则会报错")
	// num_bool, error := strconv.ParseBool("T")
	num_bool, error := strconv.ParseBool("0")
	fmt.Println(num_bool, "\terr:", error)
	fmt.Println(reflect.TypeOf(num_bool))
}
func main () {
	int_to_int() 	// 整->整
	string_to_int() // 字符串->整
	int_to_string()	// 整->字符串
	parse_num_api() // 特殊的转换api
}
