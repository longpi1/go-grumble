package user

import (
	"fmt"
	user2 "go-web/internal/repository/user"
	"go-web/internal/services/user"
)

var _ Handler = (*userHandler)(nil)

type userHandler struct {
	userService user.Service
}

type Handler interface {
	i()

	// Create 创建员工
	Create()

	// search 通过id获取员工详情
	Search()

	// Delete 删除员工
	Delete()

	// List 员工列表
	List()

	// Modify 修改员工信息
	Modify()
}

func New() Handler {
	return &userHandler{
		userService: user.New(),
	}
}

func (h *userHandler) Create() {
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
	//先用实体类存储数据
	userDate := user2.NewUserDate(name, startTime, department, position)

	h.userService.Create(&userDate)

}

func (h *userHandler) Search() {
	fmt.Println("----------查找员工----------")
	fmt.Print("请选择查找员工的编号（-1退出）：")
	//默认为-1
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		//放弃查找
		return
	}
	fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
	if !h.userService.SearchById(id) {
		fmt.Println("        用户不存在")
	}
}

func (h *userHandler) Delete() {
	fmt.Println("----------删除员工----------")
	fmt.Print("请输入你要删除员工的ID(-1退出)：")
	//默认为-1
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		//放弃删除操作
		return
	}
	if h.userService.Delete(id) {
		fmt.Println("----------删除成功----------")
	} else {
		fmt.Println("-----删除失败,ID不存在------")
	}

}

func (h *userHandler) List() {
	fmt.Println("----------列表----------")
	fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
	//按照内容过滤
	fmt.Println("员工信息类型（支持姓名、部门、职位过滤）：")
	fmt.Println("        1 全部员工信息")
	fmt.Println("        2 按姓名过滤员工信息")
	fmt.Println("        3 按部门过滤员工信息")
	fmt.Println("        4 按职位过滤员工信息")
	fmt.Println("        5 根据员工Id排序")
	fmt.Println("        6 根据员工入职日期排序")
	h.userService.List()

}

func (h *userHandler) Modify() {
	fmt.Println("----------修改员工----------")
	fmt.Print("请选择修改员工的编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
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
    //先用实体类存储数据
	userDate := user2.NewUserDate(name, startTime, department, position)

	if h.userService.Modify(id, &userDate) {
		fmt.Println("----------修改成功----------")
	} else {
		fmt.Println("----------修改失败,ID不存在----------")
	}
}

func (h *userHandler) i() {}
