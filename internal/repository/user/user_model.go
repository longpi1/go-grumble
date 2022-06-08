package user

import (
	"fmt"
	"time"
)

// User  
//go:generate gormgen -structs User -input . 
type User struct {
	Id int //客户id
	Name string //名称
	StartTime time.Time  // 入职时间
	Department string // 部门
	Position string // 职位
}

//返回客户的信息
func (this User) GetInfo() string {
	info := fmt.Sprintf("lptest%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.StartTime,
		this.Department, this.Position)
	return info
}

