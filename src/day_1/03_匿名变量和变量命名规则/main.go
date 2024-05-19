package main
import "fmt"

func get_value() (int, int ){
	return 1, 2
}

func go_set_name_format() {
	// 1_a // 不得以数字开头命名
	a_1 := "kangkang"
	_1_a := "wangming"
	fmt.Println(a_1, _1_a)
	fmt.Println("golang 不得以数字开头进行命名")
}

func main() {
	// 匿名变量
	var a, _ = 1, 2 // 加下划线是匿名变量，便是不用第二个值。不对第二个值进行特殊存储
	fmt.Println(a)
	
	// 使用场景（在获取函数返回值时，只想获取第一个值，不获取第二个值时就可以使用匿名函数）
	var x, _ = get_value()
	fmt.Println(x) 

	// 变量名规则
	go_set_name_format()
}