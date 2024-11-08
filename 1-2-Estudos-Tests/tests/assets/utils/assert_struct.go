package utils

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type StructAssert struct {
	PropertyName string
	Type         reflect.Kind
	Value        interface{}
}

func findStructAssert(propertyName string, propertiesToAssert []StructAssert) *StructAssert {
	for _, property := range propertiesToAssert {
		if property.PropertyName == propertyName {
			return &property
		}
	}

	return nil
}

func AssertStruct(t *testing.T, data interface{}, propertiesToAssert []StructAssert) {
	value := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	typeIsStruct := typ.Kind() == reflect.Struct
	require.True(t, typeIsStruct, fmt.Sprintf("Input type must be a Struct"))

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		valueField := value.Field(i)
		structAssert := findStructAssert(field.Name, propertiesToAssert)
		require.NotNil(t, structAssert, fmt.Sprintf("Not found property to assert %v", field.Name))
		require.Equal(t, field.Name, structAssert.PropertyName, fmt.Sprintf("Expected Name to be %v, but property is %v", field.Name, structAssert.PropertyName))
		require.Equal(t, valueField.Kind(), structAssert.Type, fmt.Sprintf("Expected Type to be %v, but property is %v", valueField.Kind(), structAssert.Type))
		require.Equal(t, valueField.Interface(), structAssert.Value, fmt.Sprintf("Expected Value to be %v, but property is %v", valueField.Interface(), structAssert.Value))
	}
}
