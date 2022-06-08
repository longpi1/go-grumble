package user

import (
	"go-web/internal/repository/user"
)

var _ Service = (*service)(nil)

type service struct {

}

type Service interface {
	i()
	Create( userDate  *user.UserData) ( err error)
	Modify(id int, userDate  *user.UserData)
	List() (listData []*user.User)
	Delete(id int) (flag  bool )
	Search(id int) (info user.User)
}

func New() Service {
	return &service{
	}
}

func (s *service) i() {}