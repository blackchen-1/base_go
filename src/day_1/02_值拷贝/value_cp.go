package main
import "fmt"

func main(){
	// 值拷贝的概念
	var x = 10
	var y = x // 首先这里会在栈上开辟一个空间给y,x将地址指向值拷贝一份到y的地址空间中
	x = 20
	fmt.Println(x, y)
}