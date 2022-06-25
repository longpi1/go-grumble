package employee

import (
	"fmt"
	"github.com/desertbit/grumble"
	"go-web/configs"
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
	// @title    Create 创建员工
	// @return    error   “错误信息”
	Create(c grumble.Context)  error

	// @title    search 通过id获取员工详情
	// @return    error   “错误信息”
	Search(c grumble.Context) error

	// @title    Delete 根据id删除员工
	// @return    error   “错误信息”
	Delete(c grumble.Context)  error

	// @title List 返回符合要求的员工列表
	// @return    error   “错误信息”
	List(c grumble.Context)  error

	// @title Modify 根据id修改员工信息
	// @return    error   “错误信息”
	Modify(c grumble.Context)  error
}

//创建Handler对象
func New() Handler {
	return &employeeHandler{
		employeeService: employee.New(),
	}
}

func (h *employeeHandler) Create(c grumble.Context)  error{
	name := c.Flags.String("name")
	if c.Flags.String("name") == "" {
		return fmt.Errorf("请输入员工名称")
	}
	startTime := c.Flags.String("startTime")
	department := c.Flags.String("department")
	position := c.Flags.String("position")
	//先构建员工传输对象
	employeeDate := employee2.NewEmployeeVo(name, startTime, department, position)
	h.employeeService.Create(&employeeDate)
	return nil
}

func (h *employeeHandler) Search(c grumble.Context) error{
	id := c.Flags.Int("id")
	if !h.employeeService.SearchById(id) {
		return fmt.Errorf("用户不存在:%d", id)
	}
	return nil
}

func (h *employeeHandler) Delete(c grumble.Context) error{
	id := c.Flags.Int("id")
	if !h.employeeService.Delete(id) {
		return fmt.Errorf("删除失败,ID不存在:%d", id)
	}
    return nil
}

func (h *employeeHandler) List(c grumble.Context) error{
	key := c.Flags.Int("key")
	//查看输入的key是否在查找类型范围
	isRange := key <= configs.SortByTime && key >= configs.QueryAll
	if !isRange {
		return fmt.Errorf("输入的类型错误：%d,请输入（1-6）进行查找", key)
	}
	value := c.Flags.String("value")
	h.employeeService.List(key, value)
	return nil
}

func (h *employeeHandler) Modify(c grumble.Context) error{
	id := c.Flags.Int("id")
	name := c.Flags.String("name")
	startTime := c.Flags.String("startTime")
	department := c.Flags.String("department")
	position := c.Flags.String("position")
	//先构建员工传输对象
	employeeDate := employee2.NewEmployeeVo(name, startTime, department, position)
	if !h.employeeService.Modify(id, &employeeDate) {
		return fmt.Errorf("修改失败,ID不存在:%d", id)
	}
	return nil
}


