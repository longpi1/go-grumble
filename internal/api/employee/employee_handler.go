package employee

import (
	"fmt"
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
	Create()

	// search 通过id获取员工详情
	Search(id int)

	// Delete 根据id删除员工
	Delete(id int)

	// List 返回符合要求的员工列表
	List( args []string)

	// Modify 根据id修改员工信息
	Modify(id int)
}

func New() Handler {
	return &employeeHandler{
		employeeService: employee.New(),
	}
}

func (h *employeeHandler) Create() {
	fmt.Println("----------添加员工----------")
	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("入职时间：")
	startTime := ""
	fmt.Scanln(&startTime)
	fmt.Print("部门：")
	department := ""
	fmt.Scanln(&department)
	fmt.Print("职位：")
	position := ""
	fmt.Scanln(&position)
	//先构建员工传输类
	employeeDate := employee2.NewEmployeeDate(name, startTime, department, position)

	h.employeeService.Create(&employeeDate)

}

func (h *employeeHandler) Search(id int) {
	fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
	if !h.employeeService.SearchById(id) {
		logrus.Error("----------用户不存在----------")
	}
}

func (h *employeeHandler) Delete(id int) {
	if h.employeeService.Delete(id) {
		fmt.Println("----------删除成功----------")
	} else {
		logrus.Error("----------删除失败,ID不存在----------")
	}

}

func (h *employeeHandler) List(args []string) {
	fmt.Println("----------列表----------")
	fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
	fmt.Println("员工信息类型（支持姓名、部门、职位过滤）：")
	fmt.Println("        1 全部员工信息")
	fmt.Println("        2 按姓名过滤员工信息")
	fmt.Println("        3 按部门过滤员工信息")
	fmt.Println("        4 按职位过滤员工信息")
	fmt.Println("        5 根据员工Id排序")
	fmt.Println("        6 根据员工入职日期排序")
	h.employeeService.List()
}

func (h *employeeHandler) Modify(id int) {
	fmt.Println("----------修改员工----------")
	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("入职时间：")
	startTime := ""
	fmt.Scanln(&startTime)
	fmt.Print("部门：")
	department := ""
	fmt.Scanln(&department)
	fmt.Print("职位：")
	position := ""
	fmt.Scanln(&position)
	//先构建员工传输类
	employeeDate := employee2.NewEmployeeDate(name, startTime, department, position)

	if h.employeeService.Modify(id, &employeeDate) {
		fmt.Println("----------修改成功----------")
	} else {
		logrus.Error("----------修改失败,ID不存在----------")
	}
}

func (h *employeeHandler) i() {}
