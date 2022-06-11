package user

import (
	"fmt"
)

// 员工实际存储类
type User struct {
	Id         int    //员工id
	Name       string //名称
	StartTime  string // 入职时间
	Department string // 部门
	Position   string // 职位
}

//返回员工的信息
func (user User) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t", user.Id, user.Name, user.StartTime,
		user.Department, user.Position)
	return info
}
