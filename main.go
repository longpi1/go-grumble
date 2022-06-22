package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	configs "go-web/configs"
	"go-web/internal/api/employee"
	"os"
	"strconv"
	"strings"
)

//查看基本命令
func help() {
	fmt.Println("------员工信息管理系统命令格式如下------")
	fmt.Println("        添加员工: add")
	fmt.Println("        修改员工: modify {id}")
	fmt.Println("        删除员工: del {id}")
	fmt.Println("        员工列表: list")
	fmt.Println("        查找员工：search {id}")
	fmt.Println("        帮助：help")
	fmt.Println("        退出系统：exit")

}
//运行命令行
func runCommand(commandStr string) {
	//先初始化控制器
	handler := employee.New()
	//命令行
	commandStr = strings.TrimSuffix(commandStr, "\n")
	if commandStr=="" {
		fmt.Println("命令行不能为空,请重新输入")
		return
	}
	//分割命令
	agrs := strings.Fields(commandStr)
	switch agrs[0] {
	case configs.AddEmployee:
		//创建员工
		handler.Create()
	case configs.ModifyEmployee:
		id, _ :=strconv.Atoi(agrs[1])
		//根据id更新员工信息
		handler.Modify(id)
	case configs.DeleteEmployee:
		id, _ :=strconv.Atoi(agrs[1])
		//根据id删除员工
		handler.Delete(id)
	case configs.EmployeeList:
		//遍历员工信息
		handler.List()
	case configs.SearchEmployeeById:
		id, _ :=strconv.Atoi(agrs[1])
		//根据员工id搜索员工信息
		handler.Search(id)
	case configs.Exit:
		os.Exit(0)
	default:
		help()
	}
}


func main() {
	//根据你的配置环境设置日志级别
	logrus.SetLevel(logrus.Level(configs.Get().Logger.Level))
	logrus.Info("员工信息管理系统启动")
	reader := bufio.NewReader(os.Stdin)
	for {
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			logrus.Error("命令格式出错，请重新输入")
		}
		//运行命令行
		runCommand(cmdString)
	}
}
