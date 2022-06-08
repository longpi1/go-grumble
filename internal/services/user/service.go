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
	Create(userDate *user.UserData)
	Modify(id int, userDate *user.UserData) bool
	List() (listData []user.User)
	Delete(id int) (flag bool)
	SearchById(id int) (info user.User)
}

func New() Service {
	return &service{}
}

func (s *service) i() {}
