//package main
//
//import (
//	"bufio"
//	"fmt"
//	"github.com/sirupsen/logrus"
//
//	"go-web/cmd/cmd"
//	_ "go-web/cmd/cmd"
//	"go-web/configs"
//	"os"
//	"strings"
//)
//
//func main()  {
//	//根据你的配置环境设置日志级别
//	logrus.SetLevel(logrus.Level(configs.Get().Logger.Level))
//	logrus.Info("员工信息管理系统启动")
//	reader := bufio.NewReader(os.Stdin)
//	for {
//		cmdString, err := reader.ReadString('\n')
//		if err != nil {
//			logrus.Error("命令格式出错，请重新输入")
//		}
//		//命令行
//		cmdString = strings.TrimSuffix(cmdString, "\n")
//		if cmdString=="" {
//			fmt.Println("命令行不能为空,请重新输入")
//			return
//		}
//		//分割命令
//		cmdStr := strings.Fields(cmdString)
//		if cmdStr[0]==configs.Exit {
//			return
//		}
//		//运行命令行
//		cmd.Execute()
//	}
//}


package main
import (
	"github.com/desertbit/grumble"
	"go-web/cmd/cmd"
)

func main() {

	grumble.Main(cmd.Employee)
}


