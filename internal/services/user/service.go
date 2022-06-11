package user

import (
	"go-web/internal/repository/user"
)

var _ Service = (*userService)(nil)

//定义userService，完成对user的增删查改
type userService struct {
	//用户存储的切片
	userList   []user.User
	//表示当前切片含有多少个员工，以及作为新员工的id号
	userId int
}
//定义用户service层接口
type Service interface {
	i()

	// @title    Create
	// @description   创建员工
	// @param     userDate        *user.UserData         "员工实体类"
	// @return    nil
	Create(userDate *user.UserData)

	// @title    Modify
	// @description   根据id修改员工信息
	// @param     id        int         "员工Id"
	// @param     userDate        *user.UserData         "员工实体类"
	// @return    bool   “表示是否修改成功”
	Modify(id int, userDate *user.UserData) bool

	// @title    List
	// @description   遍历员工（支持内容过滤以及排序）
	List()


	// @title    Delete
	// @description   根据id删除员工信息
	// @param     id        int         "员工Id"
	// @return    bool   “表示是否删除成功”
	Delete(id int) bool

	// @title    SearchById
	// @description   根据id搜索员工
	// @param     id        int         "员工Id"
	// @return    bool   “表示是否查找成功”
	SearchById(id int) bool
}

func New() Service {
	return &userService{}
}

func (s *userService) i() {}
