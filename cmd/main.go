package main

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"go-web/cmd/cmd"
	"go-web/configs"
	"os"
)

func main()  {
	//根据你的配置环境设置日志级别
	logrus.SetLevel(logrus.Level(configs.Get().Logger.Level))
	logrus.Info("员工信息管理系统启动")
	reader := bufio.NewReader(os.Stdin)
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			logrus.Error("命令格式出错，请重新输入")
		}
		//运行命令行
		cmd.Execute()
	}
}