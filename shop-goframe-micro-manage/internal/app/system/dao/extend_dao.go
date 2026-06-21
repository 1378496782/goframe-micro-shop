package dao

import (
	"reflect"
)

// 解决多表排除字段不能用的情况
// 例:FieldsEx(internal.SysUser,"a","ex_col1","ex_col2")
func FieldsEx(obj interface{}, preFix string, exFields ...string) []string {
	v := reflect.ValueOf(obj)
	columnsMethod := v.MethodByName("Columns")
	if !columnsMethod.IsValid() {
		return []string{}
	}
	result := columnsMethod.Call([]reflect.Value{})

	var fields []string
	for i := 0; i < result[0].NumField(); i++ {
		f := result[0].Field(i)
		if f.Kind() == reflect.String {

			if !isInArray(f.String(), exFields) {
				fields = append(fields, preFix+"."+f.String())
			}
		}
	}
	return fields
}

func isInArray(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
