package main

import "fmt"
import "strings"
/*
	练习: 实现一个学生管理系统（无结构体版本, 使用切片嵌套map实现）
	1. 学生有id、姓名、年龄、分数等信息
	2. 程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
*/

// 单个学生信息存储和多人学生存储库 
var all_stus_store_library = make([]map[int]string, 0, 100) // 存储库
var stu_counts = 0  // 当前系统学生总人数

// 输出系统选项
func print_system_mode() {
	fmt.Println(`
-----------------  学生管理系统 -------------------
		1. 增加学员信息
		2. 删除学员信息
		3. 更改学员信息
		4. 查询指定学员信息
		5. 查询所有学员信息
		6. 退出系统
----------------------------------------------------
			`)
}

// 信息输入
func input_stus_msg(id, name, age, score *string) map[int]string {
	fmt.Print("输入【学员编号】: " )
	fmt.Scanln(id)
	fmt.Print("输入【学员姓名】: ")
	fmt.Scanln(name)
	fmt.Print("输入【学员年龄】: ")
	fmt.Scanln(age)
	fmt.Print("输入【学员分数】: ")
	fmt.Scanln(score)

	var single_stu_msg = map[int]string {
		1 : *id,
		2 : *name,
		3 : *age,
		4 : *score,
	}
	return single_stu_msg
}

// 增加学员信息
func insert_stus_msg(id, name, age, score *string) {
	var is_continue string
	for true {
		// 输入信息接口
		var single_stu_msg = input_stus_msg(id, name, age, score)
		
		if  stu_counts > 50 {
			fmt.Println("人员已满,无法添加")
			break;
		}
		
		all_stus_store_library = append(all_stus_store_library, single_stu_msg)
		stu_counts ++
	
		fmt.Print("是否继续添加【Y/N】:")
		fmt.Scanln(&is_continue)
		if strings.ToUpper(is_continue) == "Y" {
			continue 
		} else {
			break
		}
	}
}

// 删除学员信息(指定学员名字)
func delete_stus_msg(store_library []map[int]string, stu_name string) {
	// 入参校验
	if store_library == nil || stu_name == "" {
		fmt.Println("store_library or stu_name is nil\n")
		return 
	}

	// 做一个空map, 用于格式化单个学员信息
	var format_single_map_msg = map[int]string{
		1 : "",
		2 : "",
		3 : "",
		4 : "", 
	}

	// 该变量用于确定是否有此学生
	var is_have_this_student int = 0

	// 遍历学员信息进行确认,并删除
	for slice_index, single_map := range store_library {
		for map_key, map_stu_msg_value := range single_map {
			if map_key == 2  && map_stu_msg_value == stu_name {
				store_library[slice_index] = format_single_map_msg
				is_have_this_student = 1
			}
		}
	}

	if (is_have_this_student == 0) {
		fmt.Print("库中没有%s信息, 请确认信息之后,再来删除\n", stu_name)
		return
	}

	stu_counts--
	fmt.Printf("删除%s信息成功\n", stu_name)
	return 
} 

// 更改一组学员信息(指定学员名字)
func update_stus_msg(store_library []map[int]string, new_stu_info map[int]string, old_stu_name string) {
	// 入参检查
	if store_library == nil || new_stu_info == nil || old_stu_name == "" {
		fmt.Println("store_library or stu_name is nil\n")
		return 
	}

	var is_have_this_student int = 0

	// 遍历学员信息进行确认并更改
	for slice_index, single_map := range store_library {
		for map_key, map_stu_msg_value := range single_map {
			if map_key == 2  && map_stu_msg_value == old_stu_name {
				store_library[slice_index] = new_stu_info // 更新新信息 
				is_have_this_student = 1
			}
		}
	}

	if (is_have_this_student == 0) {
		fmt.Print("库中没有%s信息, 请确认信息之后,再来更改\n", old_stu_name)
		return
	}

	fmt.Printf("更改%s信息成功\n", new_stu_info[2])
	return
}

// 查询指定学员信息(指定学员名字)
func select_stu_msg(store_library []map[int]string, stu_name string) {
	// 入参检查
	if store_library == nil || stu_name == "" {
		fmt.Println("store_library or stu_name is nil \n")
		return 
	}

	var is_have_this_student int = 0

	// 遍历学员信息进行确认并更改
	for slice_index, single_map := range store_library {
		for map_key, map_stu_msg_value := range single_map {
			if map_key == 2  && map_stu_msg_value == stu_name {
				fmt.Print("编号\t\t姓名\t\t年龄\t\t分数\n")
				for k:= 1; k < 5; k++ {
					fmt.Printf("%s\t\t", store_library[slice_index][k])
				}
				is_have_this_student = 1
			}
		}
	}

	if (is_have_this_student == 0) {
		fmt.Print("库中没有%s信息, 请确认信息之后,再来更改\n", stu_name)
		return
	}

	fmt.Println()
	return 
}

// 查询所有学员信息
func select_all_stus(store_library []map[int]string) {
	
	// 入参检查
	if store_library == nil {
		fmt.Println("store_library is nil")
		return 
	}

	fmt.Println("----------------------------------------------------")
	fmt.Printf("当前学员数量: %d\n", stu_counts)
	fmt.Print("编号\t\t姓名\t\t年龄\t\t分数\n")

	for _, v_map := range store_library {

		for key:= 1; key < 5; key++ {
			fmt.Printf("%s\t\t", v_map[key])
		}
		fmt.Println()
	
	}

	fmt.Println()
	fmt.Println("----------------------------------------------------")
}

func main() {

	var choose_mode int
	var break_loop = false
	var id, old_stu_name, name, age, score  string

	for true {
		// 提示系统模式
		print_system_mode()

		fmt.Print("输入【模式】: ")
		fmt.Scanln(&choose_mode)

		switch choose_mode {
			case 1 :
				// 插入信息接口
				insert_stus_msg(&id, &name, &age, &score)
			case 2 :
				// 删除学员信息
				fmt.Print("请输入要删除的名字:")
				fmt.Scanln(&old_stu_name)
				delete_stus_msg(all_stus_store_library, old_stu_name) //固定名字测试
			case 3 :
				// 更改学员信息
				fmt.Print("请输入要更改的名字:")
				fmt.Scanln(&old_stu_name)
				update_single_stu_library := input_stus_msg(&id, &name, &age, &score)
				update_stus_msg(all_stus_store_library, update_single_stu_library, old_stu_name)
			case 4 :
				// 查询学员信息(根据名字)
				fmt.Print("请输入要查询的名字:")
				fmt.Scanln(&name)
				select_stu_msg(all_stus_store_library, name)
			case 5 :
				select_all_stus(all_stus_store_library)
			default :
				// 退出系统
				break_loop = true
		}

		if break_loop {
			break
		}
	}

}