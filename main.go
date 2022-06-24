package main

import (
	"github.com/desertbit/grumble"
	"github.com/sirupsen/logrus"
	"go-web/cmd/cmd"
	"go-web/configs"
)

func main() {
	//根据配置环境设置日志级别
	logrus.SetLevel(logrus.Level(configs.Get().Logger.Level))
	logrus.Info("员工信息管理系统启动")
	grumble.Main(cmd.Employee)
}
