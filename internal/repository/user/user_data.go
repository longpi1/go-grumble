package user

//数据实体
type UserData struct {
	Id         int    //员工id
	Name       string //名称
	StartTime  string // 入职时间
	Department string // 部门
	Position   string // 职位
}

func NewUserDate(name string, startTime string, department string,
	position string) UserData {
	return UserData{
		Name:       name,
		StartTime:  startTime,
		Department: department,
		Position:   position,
	}
}
