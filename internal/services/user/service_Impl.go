package user

import (
	"fmt"
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

func (s *service) Delete(id int) bool {
	index := s.FindById(id)
	if index == -1 {
		return false
	}
	s.user = append(s.user[:s.user[index].Id-1], s.user[s.user[index].Id:]...)
	return true
}

func (s *service) Modify(id int, userDate *user.UserData) bool {
	index := s.FindById(id)
	if index == -1 {
		return false
	}
	s.user[index].Name = userDate.Name
	s.user[index].Department = userDate.Department
	s.user[index].Position = userDate.Position
	s.user[index].StartTime = userDate.StartTime
	s.user = append(append(s.user[:s.user[index].Id-1], s.user[index]), s.user[s.user[index].Id:]...)
	return true
}

func (s *service) SearchById(id int) bool {
	index := -1
	for i := 0; i < len(s.user); i++ {
		if s.user[i].Id == id {
			index = i
		}
	}
	if index != -1 {
		fmt.Println(s.user[index].GetInfo())
		return true
	}
	return false
}

func (s *service) List() (listInfo []user.User) {
	return s.user
}

//根据ID查找客户在切片中对应下标，如果没有该客户，返回-1
func (s *service) FindById(id int) int {
	index := -1
	//遍历cv.customers 切片
	for i := 0; i < len(s.user); i++ {
		if s.user[i].Id == id {
			index = i
		}
	}
	return index
}
