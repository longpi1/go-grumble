package employee

import (
	"fmt"
	"github.com/desertbit/grumble"
	"github.com/sirupsen/logrus"
	employee2 "go-web/internal/repository/employee"
	"go-web/internal/services/employee"
)

var _ Handler = (*employeeHandler)(nil)

//handler层结构体
type employeeHandler struct {
	//内置service层对象
	employeeService employee.Service
}

//定义员工handler层接口
type Handler interface {
	i()

	// Create 创建员工
	Create(c grumble.Context)

	// search 通过id获取员工详情
	Search(c grumble.Context)

	// Delete 根据id删除员工
	Delete(c grumble.Context)

	// List 返回符合要求的员工列表
	List(c grumble.Context)

	// Modify 根据id修改员工信息
	Modify(c grumble.Context)
}

func New() Handler {
	return &employeeHandler{
		employeeService: employee.New(),
	}
}

func (h *employeeHandler) Create(c grumble.Context) {
	fmt.Println("----------添加员工----------")
	name := c.Flags.String("name")
	startTime := c.Flags.String("startTime")
	department := c.Flags.String("department")
	position := c.Flags.String("position")
	//先构建员工传输对象
	employeeDate := employee2.NewEmployeeVo(name, startTime, department, position)
	h.employeeService.Create(&employeeDate)
}

func (h *employeeHandler) Search(c grumble.Context) {
	fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
	id := c.Flags.Int("id")
	if !h.employeeService.SearchById(id) {
		logrus.Error("----------用户不存在----------")
	}
}

func (h *employeeHandler) Delete(c grumble.Context) {
	id := c.Flags.Int("id")
	if h.employeeService.Delete(id) {
		fmt.Println("----------删除成功----------")
	} else {
		logrus.Error("----------删除失败,ID不存在----------")
	}

}

func (h *employeeHandler) List(c grumble.Context) {
	fmt.Println("----------员工列表----------")
	fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
	key := c.Flags.Int("key")
	value := c.Flags.String("value")
	h.employeeService.List(key, value)
}

func (h *employeeHandler) Modify(c grumble.Context) {
	fmt.Println("----------修改员工----------")
	id := c.Flags.Int("id")
	name := c.Flags.String("name")
	startTime := c.Flags.String("startTime")
	department := c.Flags.String("department")
	position := c.Flags.String("position")
	//先构建员工传输对象
	employeeDate := employee2.NewEmployeeVo(name, startTime, department, position)
	if h.employeeService.Modify(id, &employeeDate) {
		fmt.Println("----------修改成功----------")
	} else {
		logrus.Error("----------修改失败,ID不存在----------")
	}
}

func (h *employeeHandler) i() {}
