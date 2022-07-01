package employee

import (
	"fmt"
	_ "github.com/sirupsen/logrus"
	"go-web/configs"
	"go-web/internal/repository/employee"
	"sort"
)

//创建员工信息
func (s *EmployeeService) Create(employeeDate *employee.EmployeeVo) {
	user := employee.Employee{
		Id:         employeeDate.Id,
		Name:       employeeDate.Name,
		StartTime:  employeeDate.StartTime,
		Department: employeeDate.Department,
		Position:   employeeDate.Position,
	}
	//添加该员工到map集合中
	s.employeeMap[employeeDate.Id] = &user
}

//根据id删除员工信息
func (s *EmployeeService) Delete(id int) bool {
	_, ok := s.employeeMap[id]
	//返回false则表示员工不存在
	if !ok {
		return false
	}
	//实现从员工map中删除一个元素
	delete(s.employeeMap, id)
	fmt.Println("----------删除成功----------")
	return true
}

//根据id更新员工信息
func (s *EmployeeService) Modify(id int, employeeDate *employee.EmployeeVo) bool {
	e, ok := s.employeeMap[id]
	if !ok {
		return false
	}
	//将实体数据传输给存储类
	e.Name = employeeDate.Name
	e.Department = employeeDate.Department
	e.Position = employeeDate.Position
	e.StartTime = employeeDate.StartTime
	//实现从用户map中更新一个元素
	s.employeeMap[id] = e
	fmt.Println("----------修改成功----------")
	return ok
}

//根据id搜索员工信息
func (s *EmployeeService) SearchById(id int) bool {
	e, ok := s.employeeMap[id]
	if ok {
		//存在则打印员工信息
		fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
		fmt.Println(e.GetInfo())
	}
	return ok
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
	s.printEmployeeList(s.employeeMap)
}

//返回排序后所有员工信息
func (s *EmployeeService) sortByKey(key int) {
	switch key {
	//按工号排序
	case configs.SortById:
		//进行升序
		var ids []int
		for id := range s.employeeMap {
			//将键按顺序存储在slice中
			ids = append(ids, id)
		}
		sort.Ints(ids)
		for _, id := range ids {
			//打印排序后的信息
			fmt.Println(s.employeeMap[id].GetInfo())
		}
	//按入职日期排序
	case configs.SortByTime:
		var employeeList []*employee.Employee
		for _, e := range s.employeeMap {
			//将value存储在slice中
			employeeList = append(employeeList, e)
		}
		//进行升序
		sort.Slice(employeeList, func(i, j int) bool {
			return employeeList[i].StartTime < employeeList[j].StartTime
		})
		for _, e := range employeeList {
			//打印排序后的信息
			fmt.Println(e.GetInfo())
		}
	}
}

//返回过滤后的员工信息
func (s *EmployeeService) queryByKey(key int, value string) {
	switch key {
	//按姓名过滤
	case configs.QueryByName:
		for _, val := range s.employeeMap {
			if val.Name == value {
				fmt.Println(val.GetInfo())
			}
		}
	//按部门过滤
	case configs.QueryByDepartment:
		for _, val := range s.employeeMap {
			if val.Department == value {
				fmt.Println(val.GetInfo())
			}
		}
	//按职位过滤
	case configs.QueryByPosition:
		for _, val := range s.employeeMap {
			if val.Position == value {
				fmt.Println(val.GetInfo())
			}
		}
	default:
		_ = fmt.Errorf("请输入正确的类型:%d", key)
	}
}

//打印员工信息
func (s *EmployeeService) printEmployeeList(employeeMap map[int]*employee.Employee) {
	fmt.Println("----------员工列表----------")
	fmt.Println("工号\t姓名\t入职时间\t部门\t职位")
	for id := range employeeMap {
		fmt.Println(employeeMap[id].GetInfo())
	}
}
