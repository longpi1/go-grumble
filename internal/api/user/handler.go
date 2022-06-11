package user

import (
	"fmt"
	user2 "go-web/internal/repository/user"
	"go-web/internal/services/user"
)

var _ Handler = (*handler)(nil)

type handler struct {
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
	return &handler{
		userService: user.New(),
	}
}

func (h *handler) Create() {
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

	userDate := user2.NewUserDate(name, startTime, department, position)

	h.userService.Create(&userDate)
}

func (h *handler) Search() {
	fmt.Println("----------查找员工----------")
	fmt.Print("请选择查找员工的编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Printf(h.userService.SearchById(id).GetInfo())
}

func (h *handler) Delete() {
	fmt.Println("----------删除员工----------")
	fmt.Print("请输入你要删除员工的ID(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		//放弃删除操作
		return
	}
	choice := ""
	fmt.Print("确认是否删除（Y/N）:")
	for {
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" || choice == "N" || choice == "n" {
			break
		}
		fmt.Print("你的输入有误，确认是否删除（Y/N）:")
	}
	if choice == "Y" || choice == "y" {
		if h.userService.Delete(id) {
			fmt.Println("----------删除成功----------")
		} else {
			fmt.Println("-----删除失败,ID不存在------")
		}
	}

}

func (h *handler) List() {
	//获取切片中员工信息
	userList := h.userService.List()
	fmt.Println("----------列表----------")
	fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
	//对员工信息进行遍历
	for i := 0; i < len(userList); i++ {
		fmt.Println(userList[i].GetInfo())
	}

}

func (h *handler) Modify() {
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

	userDate := user2.NewUserDate(name, startTime, department, position)

	if h.userService.Modify(id, &userDate) {
		fmt.Println("----------修改成功----------")
	} else {
		fmt.Println("----------修改失败----------")
	}
}

func (h *handler) i() {}
