package reflectutils

import (
	"fmt"
	"reflect"
)

type StructField struct {
	Name  string
	Tag   string
	Value interface{}
}

type StructHelper struct {
	object    interface{}
	setFields map[string]bool
}

func NewStructHelper(object interface{}, setFields ...string) (*StructHelper, error) {
	objectType := reflect.TypeOf(object)
	if nil == objectType {
		return nil, fmt.Errorf("object is nil")
	}
	objectValue := reflect.ValueOf(object)

	if objectValue.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("object not a pointer to struct")
	}

	if objectValue.IsNil() {
		return nil, fmt.Errorf("pointer is nil")
	}

	if objectValue.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("object not a pointer to struct")
	}

	elemType := objectValue.Elem().Type()
	helper := &StructHelper{object: object, setFields: make(map[string]bool)}
	for _, field := range setFields {
		_, exist := elemType.FieldByName(field)
		if !exist {
			return nil, fmt.Errorf("field %s not found", field)
		}
		helper.setFields[field] = true
	}

	return helper, nil
}

func (helper *StructHelper) Object() interface{} {
	return helper.object
}

func (helper *StructHelper) Set(fields ...string) error {
	objectType := reflect.ValueOf(helper.object).Elem().Type()

	for _, field := range fields {
		fieldType, exist := objectType.FieldByName(field)
		if !exist {
			return fmt.Errorf("field %s not found", field)
		}
		if fieldType.Tag != "" {
			helper.setFields[field] = true
		}
	}

	return nil
}

func (helper *StructHelper) SetAll() {
	objectType := reflect.ValueOf(helper.object).Elem().Type()
	for i := 0; i < objectType.NumField(); i++ {
		fieldType := objectType.Field(i)
		if fieldType.Tag != "" {
			helper.setFields[fieldType.Name] = true
		}
	}
}

func (helper *StructHelper) Reset(fieldName string) {
	delete(helper.setFields, fieldName)
}

func (helper *StructHelper) ResetAll() {
	helper.setFields = make(map[string]bool)
}

func (helper *StructHelper) SetFields() []*StructField {
	var fields []*StructField
	objectValue := reflect.ValueOf(helper.object).Elem()
	objectType := objectValue.Type()
	for field, _ := range helper.setFields {
		fieldType, _ := objectType.FieldByName(field)
		fields = append(fields, &StructField{field, string(fieldType.Tag), objectValue.FieldByName(field).Interface()})
	}
	return fields
}

//args like argName1, argValue1, argName2, argValue2,...
func (helper *StructHelper) SetArgs() []interface{} {
	objectValue := reflect.ValueOf(helper.object).Elem()
	objectType := objectValue.Type()
	var args []interface{}
	for field, _ := range helper.setFields {
		fieldType, _ := objectType.FieldByName(field)
		args = append(args, string(fieldType.Tag))
		args = append(args, objectValue.FieldByName(field).Interface())
	}
	return args
}

func (helper *StructHelper) SetFieldLength() int {
	return len(helper.setFields)
}
