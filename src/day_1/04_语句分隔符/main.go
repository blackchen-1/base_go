package main
import "fmt"

func main() {
	// golang是使用;或者换行符 作为语句分隔符

	// ;作为语句分隔符
	var x = 100; 
	var y = 200;
	fmt.Println(x, y)

	// 换行作为语句分隔符(推荐)
	var a = 100
	var b = 200
	fmt.Println(a, b)
}