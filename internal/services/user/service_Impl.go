package user

import (
	"fmt"
	"go-web/internal/repository/user"
	"sort"
)

//创建员工信息，id自增
func (s *userService) Create(userDate *user.UserData) {
	u := user.User{
		Name:       userDate.Name,
		StartTime:  userDate.StartTime,
		Department: userDate.Department,
		Position:   userDate.Position,
	}
	s.userId++
	u.Id = s.userId
	s.user = append(s.user, u)
}

//根据id删除员工信息
func (s *userService) Delete(id int) bool {
	index := s.FindById(id)
	if index == -1 {
		return false
	}
	s.user = append(s.user[:s.user[index].Id-1], s.user[s.user[index].Id:]...)
	return true
}

//根据id更新员工信息
func (s *userService) Modify(id int, userDate *user.UserData) bool {
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

//根据id搜索员工信息
func (s *userService) SearchById(id int) bool {
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

//遍历员工信息
func (s *userService) List() {
	fmt.Print("类型：")
	kind := -1
	fmt.Scanln(&kind)
	if kind == 1 {
		//返回所有员工信息
		s.QueryAll()
	} else if kind == 5 || kind == 6 {
		//返回排序后所有员工信息
		s.SortByKey(kind)
	} else {
		fmt.Print("查找内容：")
		value := ""
		fmt.Scanln(&value)
		//返回过滤后的员工信息
		s.QueryByKey(kind, value)
	}
}

//返回所有员工信息
func (s *userService) QueryAll() {
	fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
	for i := 0; i < len(s.user); i++ {
		fmt.Println(s.user[i].GetInfo())
	}
}

//返回排序后所有员工信息
func (s *userService) SortByKey(key int) {
	ret := s.user
	switch key {
	//按工号排序
	case 5:
		//进行升序 i小于j 为升序
		sort.Slice(ret, func(i, j int) bool {
			return ret[i].Id < ret[j].Id
		})
	//按入职日期排序
	case 6:
		//进行升序
		sort.Slice(ret, func(i, j int) bool {
			return ret[i].StartTime < ret[j].StartTime
		})
	}
	fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
	for i := 0; i < len(ret); i++ {
		fmt.Println(ret[i].GetInfo())
	}
}

//返回过滤后的员工信息
func (s *userService) QueryByKey(key int, value string) {
	ret := make([]user.User, len(s.user))
	switch key {
	//按姓名过滤
	case 2:
		for _, val := range s.user {
			if val.Name == value {
				ret = append(ret, val)
			}
		}
	//按部门过滤
	case 3:
		for _, val := range s.user {
			if val.Department == value {
				ret = append(ret, val)
			}
		}
	//按职位过滤
	case 4:
		for _, val := range s.user {
			if val.Position == value {
				ret = append(ret, val)
			}
		}
	default:
		fmt.Println("请输入正确的类型 ")
	}
	fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
	for i := 1; i < len(ret); i++ {
		fmt.Println(ret[i].GetInfo())
	}
}

//根据ID查找客户在员工中对应下标，如果没有该员工，返回-1
func (s *userService) FindById(id int) int {
	index := -1
	//遍历员工 切片
	for i := 0; i < len(s.user); i++ {
		if s.user[i].Id == id {
			index = i
		}
	}
	return index
}
