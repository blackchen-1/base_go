package main

import "fmt"

/*
	1. map函数声明并初始化
	(1). 直接声明并初始化
	(2). 通过make初始化变量
*/
func map_init() {
	// 1. 直接声明并初始化一个map变量
	var m1 = map[string]string {
		"name" : "kangkang",
		"age"  : "32",
		"sex"  : "M" , // 若大括号不跟其后，需要加,结尾
	}
	fmt.Println(m1)

	// 2. 通过make函数初始化一个map变量
	var m2 = make(map[string]string)
	m2["apple"] = "sweet"
	m2["banana"] = "bad"
	m2["orange"] = "nice"
	fmt.Println(m2)
}

/*
	2. map函数使用
	map的特点：支持key查询
	(1). 增删改查
*/
func map_oprate_method() {
	// 1. 声明并初始化map变量
	var m1 = map[string]string{"name":"kangkang", "age":"30", "sex":"M"}

	// 2. 支持key查询 map对象[key]
	fmt.Println("查询操作:", m1)
	fmt.Println(m1["name"])
	fmt.Println(m1["age"])
	fmt.Println(len(m1))
	
	// 3. 判断是否存在某个键操做
	value, is_ok := m1["names"] //如果该键存在返回键对应的值和true, 不存在返回nil和false
	fmt.Println(value, is_ok)
	
	// 4. 写入一个key-value(前提是该key在map中不存在，则为创建。 否则为修改)
	m1["name_1"] = "张三"
	m1["age_1"] = "20"
	m1["sex_1"] = "M"
	fmt.Println("添加操作:", m1)

	// 5. 修改一个key-value
	m1["name"] = "李四"
	m1["age"] = "30"
	m1["sex"] = "W"
	fmt.Println("修改操作:", m1)

	// 6. 删除一个key-value
	delete(m1, "name")
	delete(m1, "age")
	delete(m1, "sex")
	fmt.Println("删除操作：", m1)
}

/*
	3. map的遍历
	(1). 使用for range遍历(map遍历出来的结果是无序的) 
	for k, v := range map对象 {

	}
*/
func map_for_range() {
	// 声明并初始化map对象m1
	var m1 = map[string]string{
		"name_1" : "张三", 
		"name_2" : "李四",
		"name_3" : "王五",
		"name_4" : "陈六",
	}

	// 通过 for range 打印map
	for k, v :=  range m1 {
		fmt.Println(k, v) // 打印出来的键值对是无序的
	}

}

// map 练习1
func map_practice() {
	// 1. map嵌套slice
	//  map类型的变量: key类型: string , value类型: []string
	var data = make(map[string][]string) // 声明并初始化一个map对象

	// (1). 添加一个省份，及其对应的多个市
	data["北京"] = []string{"朝阳", "海淀", "昌平"}
	data["海南"] = []string{"海口", "澄迈", "三亚"}
	data["河北"] = []string{"唐山", "石家庄", "保定"}
	fmt.Println(data)

	// (2). 查询海南的第一个城市
	fmt.Println(data["北京"][0])
	fmt.Println(data["海南"][0])
	fmt.Println(data["河北"][0])

	//  (3). 遍历每一个省份及对应的城市名
	for k, map_v := range data {
		fmt.Printf("所属省份: %s\n", k)
		for i, slice_v := range map_v {
			fmt.Printf("所属的第%d个城市: %s\n", i+1, slice_v)
		}
		fmt.Println()
	}
}

func main() {
	// 1. map声明并初始化
	// map_init()
	
	// 2. map使用
	// map_oprate_method()

	// 3. map的遍历
	// map_for_range()

	// 4. map练习1
	map_practice()
	
}
