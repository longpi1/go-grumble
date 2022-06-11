package user

import (
	"go-web/internal/repository/user"
)

var _ Service = (*service)(nil)

type service struct {
	user       []user.User
	customerId int
}

type Service interface {
	i()
	//创建员工
	Create(userDate *user.UserData)
	//修改员工
	Modify(id int, userDate *user.UserData) bool
	//遍历员工
	List() (listData []user.User)
	//删除员工
	Delete(id int) (flag bool)
	//根据id搜索员工
	SearchById(id int) (info user.User)
}

func New() Service {
	return &service{}
}

func (s *service) i() {}
