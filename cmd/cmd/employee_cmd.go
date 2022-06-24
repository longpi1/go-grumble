package cmd

import (
	"fmt"
	"github.com/desertbit/grumble"
	"go-web/configs"
	"go-web/internal/api/employee"
	"go-web/pkg/errors"
	"time"
)

//先初始化控制器
var handler = employee.New()

var Employee = grumble.New(&grumble.Config{
	Name:  "员工管理系统",
	Flags: func(f *grumble.Flags) {},
})

//添加的命令
var addCmd = &grumble.Command{
	Name: configs.AddEmployee,
	Help: "添加员工",
	Flags: func(f *grumble.Flags) {
		f.String("t", "startTime", time.Now().String(), "入职时间")
		f.String("n", "name", "", "名称")
		f.String("d", "department", "", "部门")
		f.String("p", "position", "", "职位")
	},
	Run: func(c *grumble.Context) error {
		if c.Flags.String("name") == "" {
			return fmt.Errorf("请输入员工名称")
		}
		handler.Create(*c)
		return nil
	},
}

//删除的命令
var delCmd = &grumble.Command{
	Name: configs.DeleteEmployee,
	Help: "按照id删除员工信息",
	Flags: func(f *grumble.Flags) {
		f.Int("i", "id", -1, "员工id")
	},
	Run: func(c *grumble.Context) error {
		if c.Flags.Int("id") == -1 {
			return fmt.Errorf("请输入员工id")
		}
		handler.Delete(*c)
		return nil
	},
}

//搜索的命令
var searchCmd = &grumble.Command{
	Name: configs.SearchEmployeeById,
	Help: "按照id搜索员工",
	Flags: func(f *grumble.Flags) {
		f.Int("i", "id", -1, "员工id")
	},
	Run: func(c *grumble.Context) error {
		if c.Flags.Int("id") == -1 {
			return fmt.Errorf("请输入员工id")
		}
		handler.Search(*c)
		return nil
	},
}

//列表的命令
var listCmd = &grumble.Command{
	Name: configs.EmployeeList,
	Help: "查看员工列表",
	Flags: func(f *grumble.Flags) {
		f.Int("k", "key", -1, "查找类型：请输入(1-6)  1 全部员工信息, 2 按姓名过滤员工信息,  3 按部门过滤员工信息,  4 按职位过滤员工信息,   5 根据员工Id排序,  6 根据员工入职日期排序")
		f.String("v", "value", "", "查找内容")
	},
	Run: func(c *grumble.Context) error {
		key := c.Flags.Int("key")
		//查看输入的key是否在查找类型范围
		flag := key <= configs.SortByTime && key >= configs.QueryAll
		if !flag {
			return fmt.Errorf("输入的类型错误：%d,请输入（1-6）进行查找", key)
		}
		handler.List(*c)
		return nil
	},
}

//修改的命令
var modifyCmd = &grumble.Command{
	Name: configs.ModifyEmployee,
	Help: "按照id修改员工信息",
	Flags: func(f *grumble.Flags) {
		f.Int("i", "id", -1, "员工id")
		f.String("t", "time", time.Now().String(), "入职时间")
		f.String("n", "name", "", "名称")
		f.String("d", "Department", "", "部门")
		f.String("p", "Position", "", "职位")
	},
	Run: func(c *grumble.Context) error {
		if c.Flags.Int("id") == -1 {
			fmt.Print("请输入员工id")
			return errors.New("请输入员工id")
		}
		handler.Modify(*c)
		return nil
	},
}

func init() {
	//初始化添加可执行的命令
	Employee.AddCommand(addCmd)
	Employee.AddCommand(delCmd)
	Employee.AddCommand(listCmd)
	Employee.AddCommand(modifyCmd)
	Employee.AddCommand(searchCmd)
}
