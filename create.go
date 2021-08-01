package thepenguin

import "reflect"

func (receiver ReflectionObject) Create() interface{}  {

	switch receiver.kind {
	case reflect.Slice:
		value :=receiver.val
		length:=value.Len()
		return reflect.MakeSlice(reflect.TypeOf(receiver.source),0, length*2)
	case reflect.Array:
		value :=receiver.val
		length:=value.Len()
		return  reflect.ArrayOf(length,reflect.TypeOf(receiver.source))
	default:
		ptr:= reflect.New(reflect.TypeOf(receiver.source))
		elem:=ptr.Elem()
		return reflect.Indirect(elem).Interface()
	}
}