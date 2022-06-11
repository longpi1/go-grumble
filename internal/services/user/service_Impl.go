package user

import (
	"fmt"
	_ "github.com/sirupsen/logrus"
	configs "go-web/config"
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
	//id自增
	s.userId++
	u.Id = s.userId
	//添加该用户到切片集合中
	s.userList = append(s.userList, u)
}

//根据id删除员工信息
func (s *userService) Delete(id int) bool {
	index := s.findById(id)
	//返回-1则表示员工不存在
	if index == -1 {
		return false
	}
	//实现从用户切片中删除一个元素
	s.userList = append(s.userList[:s.userList[index].Id-1], s.userList[s.userList[index].Id:]...)
	return true
}

//根据id更新员工信息
func (s *userService) Modify(id int, userDate *user.UserData) bool {
	index := s.findById(id)
	if index == -1 {
		//不存在则返回false
		return false
	}
	//将实体数据传输给存储类
	s.userList[index].Name = userDate.Name
	s.userList[index].Department = userDate.Department
	s.userList[index].Position = userDate.Position
	s.userList[index].StartTime = userDate.StartTime
	//实现从用户切片中更新一个元素
	s.userList = append(append(s.userList[:s.userList[index].Id-1], s.userList[index]), s.userList[s.userList[index].Id:]...)
	return true
}

//根据id搜索员工信息
func (s *userService) SearchById(id int) bool {
	index := -1
	for i := 0; i < len(s.userList); i++ {
		if s.userList[i].Id == id {
			index = i
		}
	}
	if index != -1 {
		fmt.Println(s.userList[index].GetInfo())
		return true
	}
	return false
}

//遍历员工信息
func (s *userService) List() {
	fmt.Print("类型：")
	kind := -1
	fmt.Scanln(&kind)
	if kind == -1 {
		return
	}
	if kind == configs.QueryAll {
		//返回所有员工信息
		s.queryAll()
	} else if kind == configs.SortById || kind == configs.SortByTime {
		//返回排序后所有员工信息
		s.sortByKey(kind)
	} else {
		fmt.Print("查找内容：")
		value := ""
		fmt.Scanln(&value)
		//返回过滤后的员工信息
		s.queryByKey(kind, value)
	}
}

//返回所有员工信息
func (s *userService) queryAll() {
	fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
	for i := 0; i < len(s.userList); i++ {
		fmt.Println(s.userList[i].GetInfo())
	}
}

//返回排序后所有员工信息
func (s *userService) sortByKey(key int) {
	//将目前的用户切片集合赋值给ret
	ret := s.userList
	switch key {
	//按工号排序
	case configs.SortById:
		//进行升序 i小于j 为升序
		sort.Slice(ret, func(i, j int) bool {
			return ret[i].Id < ret[j].Id
		})
	//按入职日期排序
	case configs.SortByTime:
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
func (s *userService) queryByKey(key int, value string) {
	ret := make([]user.User, len(s.userList))
	switch key {
	//按姓名过滤
	case configs.QueryByName:
		for _, val := range s.userList {
			if val.Name == value {
				ret = append(ret, val)
			}
		}
	//按部门过滤
	case configs.QueryByDepartment:
		for _, val := range s.userList {
			if val.Department == value {
				ret = append(ret, val)
			}
		}
	//按职位过滤
	case configs.QueryByPosition:
		for _, val := range s.userList {
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
func (s *userService) findById(id int) int {
	index := -1
	//遍历员工 切片
	for i := 0; i < len(s.userList); i++ {
		if s.userList[i].Id == id {
			index = i
		}
	}
	return index
}
