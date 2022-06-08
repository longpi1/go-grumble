package user

import (
	"fmt"
	user2 "go-web/internal/repository/user"
	"go-web/internal/services/user"
	"time"
)

var _ Handler = (*handler)(nil)

type handler struct {
	userService user.Service
}

type Handler interface {
	i()

	// Create 创建客户
	Create()

	// search 通过id获取客户详情
	Search()

	// Delete 删除客户
	Delete()

	// List 客户列表
	List()

	// Modify 修改客户信息
	Modify()
}


func New() Handler {
	return &handler{
		userService: user.New(),
	}
}

func (h *handler) Create() {
	fmt.Println("----------添加客户----------")


	fmt.Print("id：(不能为负数或者0)")
	id := 0
	fmt.Scanln(&id)
	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)

	fmt.Print("入职时间：")
	startTime := time.Time{}
	fmt.Scanln(&startTime)
	fmt.Print("部门：")
	department := ""
	fmt.Scanln(&department)
	fmt.Print("职位：")
	position := ""
	fmt.Scanln(&position)

	user := user2.NewUserDate(id,name,startTime,department,position)

    h.userService.Create(&user)
}

func (h *handler) Search() {
	fmt.Println("----------查找客户----------")
	fmt.Print("请选择查找客户的编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	h.userService.Search(id).GetInfo()

}

func (h *handler) Delete() {
	fmt.Println("----------删除客户----------")
	fmt.Print("请输入你要删除客户的ID(-1退出)：")
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
		//调用
		h.userService.Delete(id)
	}


}

func (h *handler) List() {
	//获取切片中客户信息
	user :=	h.userService.List()
	fmt.Println("----------列表----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	//对客户信息进行遍历
	for i := 0; i < len(user); i++ {
		fmt.Println(user[i].GetInfo())
	}

}

func (h *handler) Modify() {
	fmt.Println("----------修改客户----------")
	fmt.Print("请选择修改客户的编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)

	fmt.Print("入职时间：")
	startTime := time.Time{}
	fmt.Scanln(&startTime)
	fmt.Print("部门：")
	department := ""
	fmt.Scanln(&department)
	fmt.Print("职位：")
	position := ""
	fmt.Scanln(&position)

	user := user2.NewUserDate2(name,startTime,department,position)

	h.userService.Modify(id,&user)
}


func (h *handler) i() {}
