package main

import (
	"fmt"
	"github.com/desertbit/grumble"
	"github.com/sirupsen/logrus"
	"go-web/cmd/cmd"
	"go-web/configs"
	"net/http"
	"net/http/pprof"
	"os"
)

func main() {
	//根据配置环境设置日志级别
	logrus.SetLevel(logrus.Level(configs.Get().Logger.Level))
	logrus.Info("员工信息管理系统启动")
	// 开启pprof，监听请求
	// 开启pprof
	// 启动一个自定义mux的http服务器
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)



go func() {
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()
	//加载grumble，初始化命令行
	grumble.Main(cmd.Employee)
	ip := "0.0.0.0:6060"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}
}
