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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
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
		employeeMap map[int]*employee.Employee
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
			fields{map[int]*employee.Employee{1: {Id: 1, Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}}},
			args{id: 1},
			true,
		},
		{
			"test",
			fields{map[int]*employee.Employee{1: {Id: 1, Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}}},
			args{id: 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EmployeeService{
				employeeMap: tt.fields.employeeMap,
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
		employeeMap map[int]*employee.Employee
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
			fields: fields{map[int]*employee.Employee{1: {Id: 1, Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}}},
			args:   args{id: 1, employeeDate: &employee.EmployeeVo{Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}},
			want:   true,
		},
		{
			name:   "test",
			fields: fields{map[int]*employee.Employee{1: {Id: 1, Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}}},
			args:   args{id: 2, employeeDate: &employee.EmployeeVo{Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EmployeeService{
				employeeMap: tt.fields.employeeMap,
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
		employeeMap map[int]*employee.Employee
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
			fields: fields{map[int]*employee.Employee{1: {Id: 1, Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}}},
			args:   args{employeeDate: &employee.EmployeeVo{Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EmployeeService{
				employeeMap: tt.fields.employeeMap,
			}
			s.Create(tt.args.employeeDate)

		})
	}
}

//测试按id删除员工信息
//@param 构建员工id
//判断是否删除
func Test_employeeService_Delete(t *testing.T) {
	type fields struct {
		employeeMap map[int]*employee.Employee
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
			fields{map[int]*employee.Employee{1: {Id: 1, Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}}},
			args{id: 1},
			true,
		},
		{
			"test",
			fields{map[int]*employee.Employee{1: {Id: 1, Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}}},
			args{id: 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EmployeeService{
				employeeMap: tt.fields.employeeMap,
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
		employeeMap map[int]*employee.Employee
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
			fields{map[int]*employee.Employee{1: {Id: 1, Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}}},
			args{1, ""},
		},
		{
			"test",
			fields{map[int]*employee.Employee{1: {Id: 1, Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}}},
			args{2, "test"},
		},
		{
			"test",
			fields{map[int]*employee.Employee{1: {Id: 1, Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}}},
			args{3, "test"},
		},
		{
			"test",
			fields{map[int]*employee.Employee{1: {Id: 1, Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}}},
			args{4, "test"},
		},
		{
			"test",
			fields{map[int]*employee.Employee{1: {Id: 1, Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}}},
			args{5, ""},
		},
		{
			"test",
			fields{map[int]*employee.Employee{1: {Id: 1, Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}}},
			args{6, ""},
		},
		{
			"test",
			fields{map[int]*employee.Employee{1: {Id: 1, Name: "test", StartTime: "2022-01-22", Department: "test", Position: "test"}}},
			args{7, ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EmployeeService{
				employeeMap: tt.fields.employeeMap,
			}
			s.List(tt.args.key, tt.args.value)
		})
	}
}
