package employee

import (
	"go-web/internal/repository/employee"
	"sync"
)

var _ Service = (*employeeService)(nil)

//定义employeeService，完成对employee的增删查改
type employeeService struct {
	sync.RWMutex
	//用切片对员工信息进行存储
	employeeList []employee.Employee
	//表示当前切片含有多少个员工，以及作为新员工的id号
	employeeId int
}

//定义用户service层接口
type Service interface {
	i()

	// @title    Create
	// @description   创建员工
	// @param     employeeDate        *employee.employeeDate         "员工传输类"
	// @return    nil
	Create(employeeDate *employee.EmployeeVo)

	// @title    Modify
	// @description   根据id修改员工信息
	// @param     id        int         "员工Id"
	// @param     employeeDate        *employee.employeeDate         "员工传输类"
	// @return    bool   “表示是否修改成功”
	Modify(id int, employeeDate *employee.EmployeeVo) bool

	// @title    List
	// @param    key  查找类型 value 查找的值
	// @description   遍历员工（支持内容过滤以及排序）
	List(key int, value string)

	// @title    Delete
	// @description   根据id删除员工信息
	// @param     id        int         "员工Id"
	// @return    bool   “表示是否删除成功”
	Delete(id int) bool

	// @title    SearchById
	// @description   根据id搜索员工
	// @param     id        int         "员工Id"
	// @return    bool   “表示是否查找成功”
	SearchById(id int) bool
}

func New() Service {
	return &employeeService{}
}

func (s *employeeService) i() {}
