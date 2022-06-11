package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	configs "go-web/config"
	"go-web/internal/api/user"
)

//交互式界面函数
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
		case configs.AddUser:
			//创建员工
			handler.Create()
		case configs.ModifyUser:
			//根据id更新员工信息
			handler.Modify()
		case configs.DeleteUser:
			//根据id删除员工
			handler.Delete()
		case configs.UserList:
			//遍历员工信息
			handler.List()
		case configs.SearchUserById:
			//根据员工id搜索员工信息
			handler.Search()
		case configs.Exit:
			logrus.Info("你已经退出员工信息管理系统账号......")
			return
		default:
			fmt.Println("请从（1-6）选择正确的选项： ")
		}
	}

}

func main() {
	//根据你的配置环境设置日志级别
	logrus.SetLevel(logrus.Level(configs.Get().Logger.Level))
	logrus.Info("员工信息管理系统启动成功")
	//交互式界面
	view()
}
