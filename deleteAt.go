package thepenguin

import "reflect"

func (receiver ReflectionObject) DeleteAt(index int) interface{} {
	value :=receiver.val
	length:=value.Len()
	newSlice:=reflect.MakeSlice(reflect.TypeOf(receiver.source),0, length*2)
	for i:=0;i<length;i++{
		if index==i{
			continue
		}
		newSlice = reflect.Append(newSlice,value.Index(i))
	}

	return newSlice
}


func (receiver ReflectionObject) DeleteRange(start int, end int) interface{} {
	value :=receiver.val
	length:=value.Len()
	newSlice:=reflect.MakeSlice(reflect.TypeOf(receiver.source),0, length*2)
	for i:=0;i<length;i++{
		if i>=start && i<=end {
			continue
		}
		newSlice = reflect.Append(newSlice,value.Index(i))
	}

	return newSlice
}