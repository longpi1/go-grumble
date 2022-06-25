package employee

import (
	"go-web/internal/repository/employee"
	"reflect"
	"testing"
)
//测试创建service对象
//@param 构建初始化的service对象
//判断是否相等
func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Service
	}{
		{
			"test",
			&EmployeeService{
				nil,
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want){
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

//测试按id搜索
//@param 构建员工实体类 员工id
//判断是否相等
func Test_employeeService_SearchById(t *testing.T) {
	type fields struct {
		employeeList []employee.Employee
		employeeId   int
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"test",
			fields{[]employee.Employee{{1,"test", "2022-01-22","test", "test"}},1},
			args{id: 1},
			true,
		},
		{
			"test",
			fields{[]employee.Employee{{1,"test", "2022-01-22","test", "test"}},1},
			args{id: 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EmployeeService{
				employeeList: tt.fields.employeeList,
				employeeId:   tt.fields.employeeId,
			}
			if got := s.SearchById(tt.args.id); got != tt.want {
				t.Errorf("SearchById() = %v, want %v", got, tt.want)
			}
		})
	}
}


//测试修改员工信息
//@param 构建员工实体类 员工id
//判断是否相等
func Test_employeeService_Modify(t *testing.T) {
	type fields struct {
		employeeList []employee.Employee
		employeeId   int
	}
	type args struct {
		id           int
		employeeDate *employee.EmployeeVo
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "test",
			fields: fields{[]employee.Employee{{1,"2022-01-22","test", "test", "test"}},1},
			args:   args{id: 1, employeeDate: &employee.EmployeeVo{Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}},
			want:   true,
		},
		{
			name:   "test",
			fields: fields{[]employee.Employee{{1,"2022-01-22","test", "test", "test"}},1},
			args:   args{id: 2, employeeDate: &employee.EmployeeVo{Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EmployeeService{
				employeeList: tt.fields.employeeList,
				employeeId:   tt.fields.employeeId,
			}
			if got := s.Modify(tt.args.id, tt.args.employeeDate); got != tt.want {
				t.Errorf("Modify() = %v, want %v", got, tt.want)
			}
		})
	}
}


//测试创建员工
//@param 构建员工实体类
//判断员工是否递增
func Test_employeeService_Create(t *testing.T) {
	type fields struct {
		employeeList []employee.Employee
		employeeId   int
	}
	type args struct {
		employeeDate *employee.EmployeeVo
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test",
			fields: fields{[]employee.Employee{{1,"test","2022-01-22", "test", "test"}},1},
			args:   args{employeeDate: &employee.EmployeeVo{Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EmployeeService{
				employeeList: tt.fields.employeeList,
				employeeId:   tt.fields.employeeId,
			}
			if s.Create(tt.args.employeeDate); s.employeeId != 2 {
				t.Errorf("创建失败：%d",s.employeeId  )
			}

		})
	}
}

//测试按id删除员工信息
//@param 构建员工id
//判断是否删除
func Test_employeeService_Delete(t *testing.T) {
	type fields struct {
		employeeList []employee.Employee
		employeeId   int
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"test",
			fields{[]employee.Employee{{1,"test","2022-01-22","test", "test"}},1},
			args{id: 1},
			true,
		},
		{
			"test",
			fields{[]employee.Employee{{1,"test","2022-01-22","test", "test"}},1},
			args{id: 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EmployeeService{
				employeeList: tt.fields.employeeList,
				employeeId:   tt.fields.employeeId,
			}
			if got := s.Delete(tt.args.id); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

//测试按类型过滤员工列表
//@param 构建不同的key
//判断输出是否相同
func Test_employeeService_List(t *testing.T) {
	type fields struct {
		employeeList []employee.Employee
		employeeId   int
	}
	type args struct {
		key   int
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"test",
			fields{[]employee.Employee{{1,"test","2022-01-22", "test",
				"test"},{2,"test1","2022-01-22", "test", "test"}},2},
			args{1,""},
		},
		{
			"test",
			fields{[]employee.Employee{{1,"test","2022-01-22", "test",
				"test"},{2,"test1","2022-01-22", "test", "test"}},2},

			args{2,"test"},
		},
		{
			"test",
			fields{[]employee.Employee{{1,"test","2022-01-22", "test",
				"test"},{2,"test1","2022-01-22", "test", "test"}},2},

			args{3,"test"},
		},
		{
			"test",
			fields{[]employee.Employee{{1,"test","2022-01-22", "test",
				"test"},{2,"test1","2022-01-22", "test", "test"}},2},
			args{4,"test"},
		},
		{
			"test",
			fields{[]employee.Employee{{1,"test","2022-01-22", "test",
				"test"},{2,"test1","2022-01-22", "test", "test"}},2},
			args{5,""},
		},
		{
			"test",
			fields{[]employee.Employee{{1,"test","2022-01-22", "test",
				"test"},{2,"test1","2022-01-22", "test", "test"}},2},
			args{6,""},
		},
		{
			"test",
			fields{[]employee.Employee{{1,"test","2022-01-22", "test",
				"test"},{2,"test1","2022-01-22", "test", "test"}},2},
			args{7,""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EmployeeService{
				employeeList: tt.fields.employeeList,
				employeeId:   tt.fields.employeeId,
			}
			s.List(tt.args.key,tt.args.value)
		})
	}
}

