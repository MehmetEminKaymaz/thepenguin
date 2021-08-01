package thepenguin

import (
	"reflect"
)

type ReflectionObject struct {
	source interface{}
	val reflect.Value
	kind reflect.Kind
}

func Analyze(source interface{}) ReflectionObject{
	return ReflectionObject{
		source: source,
		val:reflect.ValueOf(source),
		kind: reflect.ValueOf(source).Kind(),
	}
}

func (receiver ReflectionObject) SetPropertyValueByIndex(index int, newValue interface{}) interface{} {
	ptr:=reflect.New(reflect.TypeOf(receiver.source))
	elem:= ptr.Elem()
	value:=receiver.val
	numberOfField := receiver.val.NumField()
	for i:=0;i<numberOfField;i++{
		if index==i{
			elem.Field(index).Set(reflect.ValueOf(newValue))
			continue
		}
		elem.Field(i).Set(value.Field(i))
	}

	return reflect.Indirect(elem).Interface()
}
func (receiver ReflectionObject) SetPropertyValueByName(name string, newValue interface{}) interface{}  {
	ptr:= reflect.New(reflect.TypeOf(receiver.source))
	elem:=ptr.Elem()
	value :=receiver.val
	numberOfField := receiver.val.NumField()
	for i:=0;i<numberOfField;i++{
		if value.Type().Field(i).Name==name{
			elem.Field(i).Set(reflect.ValueOf(newValue))
			continue
		}
		elem.Field(i).Set(value.Field(i))
	}

	return reflect.Indirect(elem).Interface()
}

func (receiver ReflectionObject) Length() int {
	value := receiver.val
	return value.Len()
}
func (receiver ReflectionObject) Cap() int  {
	value := receiver.val
	return value.Cap()
}







