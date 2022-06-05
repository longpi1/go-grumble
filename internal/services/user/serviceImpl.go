package user

import (
	"go-web/internal/repository/user"
)




var UserById = make(map[int]*user.User)
var UserByName = make(map[string][]*user.User)

func (s *service)  Create(userDate *user.UserData) (err error) {
	user := user.User{
		Id: userDate.Id,
		Name:   userDate.Name,
		StartTime: userDate.StartTime,
		Department:   userDate.Department ,
		Position:  userDate.Position,
	}
	if err := checkUserById(userDate.Id); err != nil {
		return err
	}else{
		store(&user)
	}
	return nil
}

func store(user *user.User) {
	UserById[user.Id] = user // 按照 id 存储
	UserByName[user.Name] = append(UserByName[user.Name], user) // 按照客户名存储
}

func checkUserById(id int) (err error){
	if UserById[id]!=nil {
		return err
	}else {
		return nil
	}
}


func (s *service)  Delete(id int) (flag bool){
	delete(UserById,id)
	return flag
}

func (s *service)  Modify(id int,userDate *user.UserData) {
	user :=UserById[id]
	user.Name = userDate.Name
	user.Department=userDate.Department
	user.Position=userDate.Position
	user.StartTime=userDate.StartTime
    store(user)
}


func (s *service)  Search(id int) (info *user.User) {
    if(id >= 0){
    	return UserById[id]
	}

	return nil
}

func (s *service)  List() (listInfo []*user.User) {
	listInfo = make([]*user.User, len(UserById))
	return listInfo
}


