package cmd

import (
	"fmt"
	"go-web/configs"
	"go-web/internal/api/employee"
	"time"

	"github.com/desertbit/grumble"
)
//先初始化控制器
var handler = employee.New()

var Employee = grumble.New(&grumble.Config{
	Name:        "员工管理系统",
	Flags:       func(f *grumble.Flags) {},
})

var addCmd = &grumble.Command{
	Name:   configs.AddEmployee,
	Help: "添加员工",
	Flags: func(f *grumble.Flags) {
		f.String("t", "time", time.Now().String(), "start time")
		f.String("n", "name", "", "name")
		f.String("d", "Department", "", "Department")
		f.String("p", "Position", "", "Position")
	},
	Run:  func(c *grumble.Context) error{
		 err := handler.Create(*c)
		 return err
	},
}


var delCmd = &grumble.Command{
	Name:   configs.DeleteEmployee,
	Help: "按照id修改员工信息",
	Run:  func(c *grumble.Context) error{
		 handler.Delete(*c)
		return nil
	},
}


var searchCmd = &grumble.Command{
	Name:   configs.SearchEmployeeById,
	Help: "按照id搜索员工",
	Run:  func(c *grumble.Context) error{
		handler.Search(*c)

		return nil
	},
}


var listCmd = &grumble.Command{
	Name:   configs.EmployeeList,
	Help: "查看员工列表",
	Run:  func(c *grumble.Context) error{
		if true {
			fmt.Println("from Earth...")
		}
		handler.List(*c)
		return nil
	},
}


var modifyCmd = &grumble.Command{
	Name:   configs.ModifyEmployee,
	Help: "按照id修改员工信息",
	Run:  func(c *grumble.Context) error{
		if true {
			fmt.Println("from Earth...")
		}
		handler.Modify(*c)
		return nil
	},
}




func init() {
	Employee.AddCommand(addCmd)
	Employee.AddCommand(delCmd)
	Employee.AddCommand(listCmd)
	Employee.AddCommand(modifyCmd)
	Employee.AddCommand(searchCmd)
}