package gocvert

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

// FieldOption 默认值填充
type FieldOption struct {
	Tag      string // name of data item
	NewField any    // new value of data item
}

// SwapOption 数据结构转换
type SwapOption struct {
	Tag    string // name of data item
	NewTag string // tag of new data item
	//NewStruct string // new data structure
}

func RewriteFields(data any, opts []FieldOption) error {

	value := reflect.ValueOf(data)
	// 必须是结构体指针类型数据
	if err := isValid(value); err != nil {
		return errors.New("data must be pointer to struct")
	}

	for _, opt := range opts {
		err := RewriteField(data, opt)
		if err != nil {
			log.Panicln(err)
		}
	}

	return nil
}

func RewriteField(data any, opt FieldOption) error {

	value := reflect.ValueOf(data)
	// 必须是结构体指针类型数据
	if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct {
		return errors.New("data must be  pointer to struct")
	}

	value = value.Elem()
	field := value.FieldByName(opt.Tag)

	if field.IsValid() && field.CanSet() {
		newVal := reflect.ValueOf(opt.NewField)
		if field.Type() == newVal.Type() {
			field.Set(newVal)
			return nil
		}
	}

	return fmt.Errorf("convert value of struct field with tag %s ", opt.Tag)
}

func SwapWithTags(data, target any, opt []SwapOption) error {

	value := reflect.ValueOf(data)
	// 必须是结构体指针类型数据
	if err := isValid(value); err != nil {
		return errors.New("data must be pointer to struct")
	}
	tValue := reflect.ValueOf(target)
	if err := isValid(tValue); err != nil {
		return errors.New("target must be pointer to struct")
	}

	var err error
	for _, op := range opt {
		err = SwapWithTag(data, target, op)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func SwapWithTag(data, target any, opt SwapOption) error {

	value := reflect.ValueOf(data)
	// 必须是结构体指针类型数据
	if err := isValid(value); err != nil {
		return errors.New("data must be pointer to struct")
	}
	tValue := reflect.ValueOf(target)
	if err := isValid(tValue); err != nil {
		return errors.New("target must be pointer to struct")
	}

	field := value.Elem().FieldByName(opt.Tag)
	tField := tValue.Elem().FieldByName(opt.NewTag)

	if field.IsValid() && tField.IsValid() && tField.CanSet() {
		if tField.Type() == field.Type() {
			tField.Set(field)
			return nil
		}
	}

	return fmt.Errorf("tag %s failed to convert value to new struct", opt.Tag)
}

func isValid(value reflect.Value) error {
	if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct {
		return errors.New("data must be  pointer to struct")
	}
	return nil
}
