package employee

//传输的数据实体
type EmployeeDate struct {
	Name       string //名称
	StartTime  string // 入职时间
	Department string // 部门
	Position   string // 职位
}
//构建传输实体
func NewEmployeeDate(name string, startTime string, department string,
	position string) EmployeeDate {
	return EmployeeDate{
		Name:       name,
		StartTime:  startTime,
		Department: department,
		Position:   position,
	}
}
