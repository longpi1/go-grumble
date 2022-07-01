package employee

import (
	"github.com/desertbit/grumble"
	"go-web/internal/services/employee"
	"reflect"
	"testing"
)

//测试创建handler对象
//@param 构建初始化的handler对象
//判断是否相等
func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Handler
	}{
		{
			"test",
			&employeeHandler{
				//构建service对象
				employeeService: employee.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//判断是否相等
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

//测试添加员工的指令
//@param 构建添加员工信息指令的flags
//判断是否添加成功
func Test_employeeHandler_Create(t *testing.T) {
	type fields struct {
		employeeService employee.Service
	}

	type args struct {
		c grumble.Context
	}
	//构建对应指令flag
	var m = map[string]*grumble.FlagMapItem{}
	m["id"] = &grumble.FlagMapItem{Value: 1, IsDefault: true}
	m["name"] = &grumble.FlagMapItem{Value: "test", IsDefault: true}
	m["startTime"] = &grumble.FlagMapItem{Value: "test", IsDefault: true}
	m["department"] = &grumble.FlagMapItem{Value: "department", IsDefault: true}
	m["position"] = &grumble.FlagMapItem{Value: "test", IsDefault: true}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				&employee.EmployeeService{},
			},
			args: args{
				c: grumble.Context{
					Flags:   m,
					Command: &grumble.Command{},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &employeeHandler{
				employeeService: tt.fields.employeeService,
			}
			//判断是否成功
			if err := h.Create(tt.args.c); (err != nil) == tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//测试按id删除员工的指令
//@param 构建删除员工信息指令的flags
//判断是否删除成功
func Test_employeeHandler_Delete(t *testing.T) {
	type fields struct {
		employeeService employee.Service
	}
	type args struct {
		c grumble.Context
	}
	//构建对应指令flag
	var m = map[string]*grumble.FlagMapItem{}
	m["id"] = &grumble.FlagMapItem{Value: 1, IsDefault: true}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				&employee.EmployeeService{},
			},
			args: args{
				c: grumble.Context{
					Flags:   m,
					Command: &grumble.Command{},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &employeeHandler{
				employeeService: tt.fields.employeeService,
			}
			//判断是否成功
			if err := h.Delete(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//测试过滤员工列表的指令
//@param 构建过滤员工信息指令的flags
func Test_employeeHandler_List(t *testing.T) {
	type fields struct {
		employeeService employee.Service
	}
	type args struct {
		c grumble.Context
	}
	//构建对应指令flag
	var m = map[string]*grumble.FlagMapItem{}
	m["key"] = &grumble.FlagMapItem{Value: 1, IsDefault: true}
	m["value"] = &grumble.FlagMapItem{Value: "", IsDefault: true}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				&employee.EmployeeService{},
			},
			args: args{
				c: grumble.Context{
					Flags:   m,
					Command: &grumble.Command{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &employeeHandler{
				employeeService: tt.fields.employeeService,
			}
			//判断是否成功
			if err := h.List(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//测试按id修改
//@param 构建修改员工信息指令的flags
//判断是否修改成功
func Test_employeeHandler_Modify(t *testing.T) {
	type fields struct {
		employeeService employee.Service
	}
	type args struct {
		c grumble.Context
	}
	//构建对应指令flag
	var m = map[string]*grumble.FlagMapItem{}
	m["id"] = &grumble.FlagMapItem{Value: 1, IsDefault: true}
	m["name"] = &grumble.FlagMapItem{Value: "test", IsDefault: true}
	m["startTime"] = &grumble.FlagMapItem{Value: "test", IsDefault: true}
	m["department"] = &grumble.FlagMapItem{Value: "department", IsDefault: true}
	m["position"] = &grumble.FlagMapItem{Value: "test", IsDefault: true}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				&employee.EmployeeService{},
			},
			args: args{
				c: grumble.Context{
					Flags:   m,
					Command: &grumble.Command{},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &employeeHandler{
				employeeService: tt.fields.employeeService,
			}
			//判断是否成功
			if err := h.Modify(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Modify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//测试按id搜索
//@param 构建搜索指令的flags
//判断是否符合
func Test_employeeHandler_Search(t *testing.T) {
	type fields struct {
		employeeService employee.Service
	}
	type args struct {
		c grumble.Context
	}
	//构建对应指令flag
	var m = map[string]*grumble.FlagMapItem{}
	m["id"] = &grumble.FlagMapItem{Value: 1, IsDefault: true}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				&employee.EmployeeService{},
			},
			args: args{
				c: grumble.Context{
					Flags:   m,
					Command: &grumble.Command{},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &employeeHandler{
				employeeService: tt.fields.employeeService,
			}
			//判断是否成功
			if err := h.Search(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
