package  main

import (
	"fmt"
	"go-web/internal/api/user"
)




func veiw(){
    handler :=user.New()
	for {
		fmt.Println("------客户信息管理系统------")
		fmt.Println("        1 添加客户")
		fmt.Println("        2 修改客户")
		fmt.Println("        3 删除客户")
		fmt.Println("        4 客户列表")
		fmt.Println("        5 查找客户")
		fmt.Println("        6 退出软件")
		fmt.Print("请选择（1-5）..")
		key :=""
		fmt.Scanln(&key)
		switch key {
		case "1":
			handler.Create()
		case "2":
			handler.Modify()
		case "3":
			handler.Delete()
		case "4":
			handler.List()
		case "5":
			handler.Search()
		case "6":
			break
		default:
			fmt.Println("请从（1-5）选择正确的选项..")
		}
	}
	fmt.Println("你已经退出客户信息管理系统账号..")
}

func main() {
	fmt.Println("客户信息管理系统开发..")
   veiw()
}
