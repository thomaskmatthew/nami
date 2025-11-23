package sql

import (
	"fmt"
	"reflect"
	"strings"
)

func CreateFieldList[T any](fields T) string {
	var fieldsList []string
	t := reflect.TypeOf(fields)
	// derefence the pointer
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		var name string
		var colType string
		name = t.Field(i).Tag.Get("sql")
		colType += " " + t.Field(i).Tag.Get("sql")
		fieldsList = append(fieldsList, fmt.Sprintf("%s %s", name, colType))
	}
	return strings.Join(fieldsList, ", ")
}

func getSelectFields[T any](fields T) string {
	var fieldsList []string

	t := reflect.TypeOf(fields)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		var name string
		name = t.Field(i).Tag.Get("sql")
		fieldsList = append(fieldsList, name)
	}

	return strings.Join(fieldsList, ", ")
}

func getValues[T any](fields T) []any {
	v := reflect.ValueOf(fields)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	args := []any{}

	for i := 0; i < t.NumField(); i++ {
		args = append(args, v.Field(i).Interface())
	}

	return args
}
