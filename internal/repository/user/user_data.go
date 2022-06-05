package user

import (
	"time"
)
//数据实体
type UserData struct {
	Id int //客户id
	Name string //名称
	StartTime time.Time  // 入职时间
	Department string // 部门
	Position string // 职位
}

func NewUserDate(id int,name string, startTime time.Time  ,department string,
	position string) UserData {
	return UserData{
		Id: id,
		Name:   name,
		StartTime: startTime,
		Department:    department,
		Position:  position,
	}
}
//除了id
func NewUserDate2(name string, startTime time.Time  ,department string,
	position string) UserData {
	return UserData{
		Name:   name,
		StartTime: startTime,
		Department:    department,
		Position:  position,
	}
}


