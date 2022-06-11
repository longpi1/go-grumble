package user

import (
	"go-web/internal/repository/user"
)

var _ Service = (*userService)(nil)

type userService struct {
	user   []user.User
	userId int
}

type Service interface {
	i()
	//创建员工
	Create(userDate *user.UserData)
	//修改员工
	Modify(id int, userDate *user.UserData) bool
	//遍历员工
	List()
	//删除员工
	Delete(id int) bool
	//根据id搜索员工
	SearchById(id int) bool
}

func New() Service {
	return &userService{}
}

func (s *userService) i() {}
