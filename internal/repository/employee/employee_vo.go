package employee

//传输的数据实体
type EmployeeVo struct {
	Name       string //名称
	StartTime  string // 入职时间
	Department string // 部门
	Position   string // 职位
}

//构建传输实体
func NewEmployeeVo(name string, startTime string, department string,
	position string) EmployeeVo {
	return EmployeeVo{
		Name:       name,
		StartTime:  startTime,
		Department: department,
		Position:   position,
	}
}
