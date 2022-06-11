package main

import (
	"fmt"
	"go-web/internal/api/user"
)

func view() {
	//先初始化控制器
	handler := user.New()
	for {
		fmt.Println("------员工信息管理系统------")
		fmt.Println("        1 添加员工")
		fmt.Println("        2 修改员工")
		fmt.Println("        3 删除员工")
		fmt.Println("        4 员工列表")
		fmt.Println("        5 查找员工")
		fmt.Println("        6 退出系统")
		fmt.Print("请选择（1-6）..")
		key := ""
		fmt.Scanln(&key)
		switch key {
		case "1":
			//创建员工
			handler.Create()
		case "2":
			//更新员工信息
			handler.Modify()
		case "3":
			handler.Delete()
		case "4":
			handler.List()
		case "5":
			//根据员工id搜索员工信息
			handler.Search()
		case "6":
			fmt.Println("你已经退出员工信息管理系统账号......")
			return
		default:
			fmt.Println("请从（1-6）选择正确的选项： ")
		}
	}

}

func main() {
	fmt.Println("员工信息管理系统")
	view()
}
