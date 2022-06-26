package employee

import (
	"fmt"
	_ "github.com/sirupsen/logrus"
	"go-web/configs"
	"go-web/internal/repository/employee"
	"sort"
)

//创建员工信息，id自增
func (s *EmployeeService) Create(employeeDate *employee.EmployeeVo) {
	user := employee.Employee{
		Name:       employeeDate.Name,
		StartTime:  employeeDate.StartTime,
		Department: employeeDate.Department,
		Position:   employeeDate.Position,
	}
	//id自增
	s.employeeId++
	user.Id = s.employeeId
	//添加该用户到切片集合中
	s.employeeList = append(s.employeeList, user)
}

//根据id删除员工信息
func (s *EmployeeService) Delete(id int) bool {
	index := s.findById(id)
	//返回-1则表示员工不存在
	if index == -1 {
		return false
	}
	//实现从用户切片中删除一个元素
	s.employeeList = append(s.employeeList[:s.employeeList[index].Id-1], s.employeeList[s.employeeList[index].Id:]...)
	fmt.Println("----------删除成功----------")
	return true
}

//根据id更新员工信息
func (s *EmployeeService) Modify(id int, employeeDate *employee.EmployeeVo) bool {
	index := s.findById(id)
	if index == -1 {
		//index为-1则表示不存在，返回false
		return false
	}
	//将实体数据传输给存储类
	s.employeeList[index].Name = employeeDate.Name
	s.employeeList[index].Department = employeeDate.Department
	s.employeeList[index].Position = employeeDate.Position
	s.employeeList[index].StartTime = employeeDate.StartTime
	//实现从用户切片中更新一个元素
	s.employeeList = append(append(s.employeeList[:s.employeeList[index].Id-1], s.employeeList[index]), s.employeeList[s.employeeList[index].Id:]...)
	fmt.Println("----------修改成功----------")
	return true
}

//根据id搜索员工信息
func (s *EmployeeService) SearchById(id int) bool {
	//默认为-1
	index := -1
	for i := 0; i < len(s.employeeList); i++ {
		if s.employeeList[i].Id == id {
			index = i
		}
	}
	if index != -1 {
		//存在则打印员工信息，返回true
		fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
		fmt.Println(s.employeeList[index].GetInfo())
		return true
	}
	return false
}

//遍历员工信息
func (s *EmployeeService) List(key int, value string) {
	if key == configs.QueryAll {
		//返回所有员工信息
		s.queryAll()
	} else if key == configs.SortById || key == configs.SortByTime {
		//返回排序后所有员工信息
		s.sortByKey(key)
	} else {
		//返回按搜索内容过滤后的员工信息
		s.queryByKey(key, value)
	}
}

//返回所有员工信息
func (s *EmployeeService) queryAll() {
	//打印员工信息
	s.printEmployeeList(s.employeeList)
}

//返回排序后所有员工信息
func (s *EmployeeService) sortByKey(key int) {
	//将目前的用户切片集合赋值给ret
	ret := s.employeeList
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
	//打印员工信息
	s.printEmployeeList(ret)
}

//返回过滤后的员工信息
func (s *EmployeeService) queryByKey(key int, value string) {
	var ret []employee.Employee
	switch key {
	//按姓名过滤
	case configs.QueryByName:
		for _, val := range s.employeeList {
			if val.Name == value {
				ret = append(ret, val)
			}
		}
	//按部门过滤
	case configs.QueryByDepartment:
		for _, val := range s.employeeList {
			if val.Department == value {
				ret = append(ret, val)
			}
		}
	//按职位过滤
	case configs.QueryByPosition:
		for _, val := range s.employeeList {
			if val.Position == value {
				ret = append(ret, val)
			}
		}
	default:
		_ = fmt.Errorf("请输入正确的类型:%d", key)
	}
	//打印员工信息
	s.printEmployeeList(ret)
}

//根据ID查找客户在员工中对应下标，如果没有该员工，返回-1
func (s *EmployeeService) findById(id int) int {
	index := -1
	//遍历员工 切片
	for i := 0; i < len(s.employeeList); i++ {
		if s.employeeList[i].Id == id {
			index = i
		}
	}
	//返回切片的index
	return index
}

//打印员工信息
func (s *EmployeeService) printEmployeeList(employeeList []employee.Employee){
	fmt.Println("----------员工列表----------")
	fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
	for i := 0; i < len(employeeList); i++ {
		fmt.Println(employeeList[i].GetInfo())
	}
}