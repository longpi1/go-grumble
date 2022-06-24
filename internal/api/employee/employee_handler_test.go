package employee

import (
	"go-web/internal/services/employee"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_employeeHandler_Create(t *testing.T) {
	type fields struct {
		employeeService employee.Service
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &employeeHandler{
				employeeService: tt.fields.employeeService,
			}
		})
	}
}

func Test_employeeHandler_Delete(t *testing.T) {
	type fields struct {
		employeeService employee.Service
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &employeeHandler{
				employeeService: tt.fields.employeeService,
			}
		})
	}
}

func Test_employeeHandler_List(t *testing.T) {
	type fields struct {
		employeeService employee.Service
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &employeeHandler{
				employeeService: tt.fields.employeeService,
			}
		})
	}
}

func Test_employeeHandler_Modify(t *testing.T) {
	type fields struct {
		employeeService employee.Service
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &employeeHandler{
				employeeService: tt.fields.employeeService,
			}
		})
	}
}

func Test_employeeHandler_Search(t *testing.T) {
	type fields struct {
		employeeService employee.Service
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &employeeHandler{
				employeeService: tt.fields.employeeService,
			}
		})
	}
}

func Test_employeeHandler_i(t *testing.T) {
	type fields struct {
		employeeService employee.Service
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &employeeHandler{
				employeeService: tt.fields.employeeService,
			}
		})
	}
}
