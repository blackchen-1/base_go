package main
import "fmt"
import "reflect"

// 内存单位转换的理解
func memory_size() {
	fmt.Println(
		"1MB = 1024KB\n",
		"1KB = 1024B\n",
		"1B = 8 bit\n",
	)
}

// 整型概念
func value_int() {
	// 整型
	fmt.Println(
		"int 整形\n",
		"理解常说的256 :\n",
		"int8 = 1Byte = 8bit = 2的8次方 = 256; 范围[-128, 127]\n",
		"uint8 = 1Byte = 8bit = 256 范围[0, 255](无符号整数从0起)\n",
		"理解常说的65535 :\n",
		"uint16 = 2Byte = 16bit = 2的16次方 = 65535(范围： 0-65535)\n",
	)
}

// 浮点型概念
func value_float() {
	fmt.Println("float 浮点型\n")
	// 浮点型
	var f1 float32 = 3.123456789123456789
	var f2 float64 = 3.123456789123456789
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println("float64对小数点后面的数取值,要比float32精准\n")

}

// 布尔类型
func value_bool() {
	var b bool
	b = true
	b = false
	fmt.Println(b)
	c := 2 > 1
	fmt.Println(c, reflect.TypeOf(c)) // 第二个输出参数是获取类型的函数
}

// 字符串类型 (字符串在内存中是一段连续存储空间)
func value_string() {

	var s string // 默认为 " "
	s = "hello world"

	/* 字符串操作*/
	// 1. 索引取字符串
	fmt.Println(string(s[1])) // 单个字符取值时,如果不用string进行字符串转义，输出的是对应的索引值的ASCLL码值
	fmt.Println(string(s[4])) 

	// 2. 切片与缺省(切片取值遵循的是，左闭右开原则)
	fmt.Println(s[0 : 4]) // 这里按道理是取到o的，但是因为切片的取值是左闭右开，所以只取了hell
	// 2.1 缺省切片取值
	fmt.Println(s[ : 4]) // ：左边没有值默认从索引零开始
	fmt.Println(s[0 : ]) // ：右边没有值默认取到字符串的最后索引
	fmt.Println(s[ : ]) // 从头取到尾

	// 3. 转义符号 \ 
	fmt.Println("hello world\n") // \n是换行符
	fmt.Println("D:\\next\\go.exe") // 要打印类似windowws系统的文件路径,需要用到转义符,对路径分隔符再进行一次转义(否则像\next中\n会误以为是换行符)
	fmt.Println("this is my name \"kangkang\"") // 在go中双引号是由特殊意义的，所以加个转义符，进行转义,则可以打印出转义符号

	// 4. 多行打印(可以用到反引号``, 在其中可以打印任何值)
	info := `
		姓名: 
		班级:
		年龄:
	`
	fmt.Println(info)
}

func main() {
	
	// 内存单位的基本了解
	memory_size()

	// int整型理解
	value_int()

	// float浮点型理解
	value_float()

	// bool布尔类型
	value_bool()

	// string 类型
	value_string()
}
