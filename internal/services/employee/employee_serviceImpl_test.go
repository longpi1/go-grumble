package employee

import (
	"go-web/internal/repository/employee"
	"testing"
)

func Test_employeeService_Create(t *testing.T) {
	type fields struct {
		employeeList []employee.Employee
		employeeId   int
	}
	type args struct {
		employeeDate *employee.EmployeeDate
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &employeeService{
				employeeList: tt.fields.employeeList,
				employeeId:   tt.fields.employeeId,
			}
		})
	}
}

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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &employeeService{
				employeeList: tt.fields.employeeList,
				employeeId:   tt.fields.employeeId,
			}
			if got := s.Delete(tt.args.id); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_employeeService_List(t *testing.T) {
	type fields struct {
		employeeList []employee.Employee
		employeeId   int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &employeeService{
				employeeList: tt.fields.employeeList,
				employeeId:   tt.fields.employeeId,
			}
		})
	}
}

func Test_employeeService_Modify(t *testing.T) {
	type fields struct {
		employeeList []employee.Employee
		employeeId   int
	}
	type args struct {
		id           int
		employeeDate *employee.EmployeeDate
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &employeeService{
				employeeList: tt.fields.employeeList,
				employeeId:   tt.fields.employeeId,
			}
			if got := s.Modify(tt.args.id, tt.args.employeeDate); got != tt.want {
				t.Errorf("Modify() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &employeeService{
				employeeList: tt.fields.employeeList,
				employeeId:   tt.fields.employeeId,
			}
			if got := s.SearchById(tt.args.id); got != tt.want {
				t.Errorf("SearchById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_employeeService_findById(t *testing.T) {
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
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &employeeService{
				employeeList: tt.fields.employeeList,
				employeeId:   tt.fields.employeeId,
			}
			if got := s.findById(tt.args.id); got != tt.want {
				t.Errorf("findById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_employeeService_queryAll(t *testing.T) {
	type fields struct {
		employeeList []employee.Employee
		employeeId   int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &employeeService{
				employeeList: tt.fields.employeeList,
				employeeId:   tt.fields.employeeId,
			}
		})
	}
}

func Test_employeeService_queryByKey(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &employeeService{
				employeeList: tt.fields.employeeList,
				employeeId:   tt.fields.employeeId,
			}
		})
	}
}

func Test_employeeService_sortByKey(t *testing.T) {
	type fields struct {
		employeeList []employee.Employee
		employeeId   int
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &employeeService{
				employeeList: tt.fields.employeeList,
				employeeId:   tt.fields.employeeId,
			}
		})
	}
}