package user

import (
	"go-web/internal/repository/user"
)

func (s *service) Create(userDate *user.UserData) {
	u := user.User{
		Name:       userDate.Name,
		StartTime:  userDate.StartTime,
		Department: userDate.Department,
		Position:   userDate.Position,
	}
	s.customerId++
	u.Id = s.customerId
	s.user = append(s.user, u)
}

func (s *service) Delete(id int) (flag bool) {
	employees := s.SearchById(id)
	if (employees == user.User{}) {
		return false
	}
	s.user = append(s.user[:employees.Id-1], s.user[employees.Id:]...)
	return true
}

func (s *service) Modify(id int, userDate *user.UserData) bool {
	employees := s.SearchById(id)
	if (employees == user.User{}) {
		return false
	}
	employees.Name = userDate.Name
	employees.Department = userDate.Department
	employees.Position = userDate.Position
	employees.StartTime = userDate.StartTime
	s.user = append(append(s.user[:employees.Id-1], employees), s.user[employees.Id:]...)
	return true
}

func (s *service) SearchById(id int) (info user.User) {
	index := -1
	for i := 0; i < len(s.user); i++ {
		if s.user[i].Id == id {
			index = i
		}
	}
	if index != -1 {
		return s.user[index]
	}
	return
}

func (s *service) List() (listInfo []user.User) {
	return s.user
}
