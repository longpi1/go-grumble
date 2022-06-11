package employee

import (
	"fmt"
)

// 员工实际存储类
type Employee struct {
	Id         int    //员工id
	Name       string //名称
	StartTime  string // 入职时间
	Department string // 部门
	Position   string // 职位
}

//返回员工的信息
func (employee Employee) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t", employee.Id, employee.Name, employee.StartTime,
		employee.Department, employee.Position)
	return info
}
