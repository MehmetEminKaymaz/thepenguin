package thepenguin

import "reflect"

func (receiver ReflectionObject) InsertAt(index int, item interface{}) interface{} {
	value:=receiver.val
	length:=value.Len()
	newSlice:=reflect.MakeSlice(reflect.TypeOf(receiver.source),0, length*2)

	for i:=0;i<length;i++{
		if index==i{
			newSlice = reflect.Append(newSlice, reflect.ValueOf(item))
			newSlice = reflect.Append(newSlice,value.Index(i))
			continue
		}
		newSlice = reflect.Append(newSlice,value.Index(i))
	}

	return newSlice
}


func (receiver ReflectionObject) Append(item interface{}) interface{} {
	value:=receiver.val
	return reflect.Append(value,reflect.ValueOf(item))
}