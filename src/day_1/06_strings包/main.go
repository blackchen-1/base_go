package main
import "fmt"
import "strings"

// strings包内API使用
func strings_api_use() {
	// 1. 将字符串转换成大小写
	var name = "Kangkang"
	var upperName = strings.ToUpper(name)
	fmt.Println(upperName)
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.ToLower(name))

	// 2. 将字符串判断前后缀, 及其中字符串是否存在
	var s = "rain yuan alvin"
	fmt.Println(strings.HasPrefix(s, "rain")) // 判断字符串前缀是否是rain
	fmt.Println(strings.HasSuffix(s, "in")) // 判断字符串尾缀是否存在
	fmt.Println(strings.Contains(s, "yuan 1")) // 判断其中是否存在yuan 1字符串

	// 3. 去除字符串两端匹配的内容
	username := "  yuan "
	// fmt.Println(strings.Trim(username, " ")) //去除首尾的匹配符(这里是去除空格)
	// username = strings.TrimSpace(username) // 去除两端空格
	username = strings.TrimLeft(username, " ") // 只去除左边的空格
	fmt.Println(username)
	fmt.Println(username == "yuan")

	// 4. index : 索引
	var s2 = "rain yuan alvin"
	fmt.Println(strings.Index(s2, "alex")) // 这个alex字符串不存在， 返回-1
	fmt.Println(strings.Index(s2, "yuan")) // 这个yuan存在, 返回字符串最开头的索引, 5
	fmt.Println(strings.Index(s2, "alvin"))// 这个alvin存在， 返回字符串最开头的索引, 10

	// 5. 分割、拼接
	var s3 = "rain yaun alvin"
	// 5.1 分割
	nameSlice := strings.Split(s3, " ") // 将字符串，以空格作为分割符。返回一个切片类型变量
	fmt.Println(nameSlice)
	fmt.Println(nameSlice[0])
	fmt.Println(nameSlice[1])
	fmt.Println(nameSlice[2])

	// 5.2 拼接
	var newStr = strings.Join(nameSlice, ",") // 以逗号为拼接符号，将切片中的内容，进行拼接。并返回新的字符串
	fmt.Println(newStr)
}

func main() {
	strings_api_use()
}