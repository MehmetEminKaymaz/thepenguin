package thepenguin

import "reflect"

func (receiver ReflectionObject) Copy() interface{}{

	switch receiver.kind {
	case reflect.Slice:
		value :=receiver.val
		length:=value.Len()
		newSlice:=reflect.MakeSlice(reflect.TypeOf(receiver.source),0, length*2)
		for i:=0;i<length;i++{
			newSlice = reflect.Append(newSlice,value.Index(i))
		}
		return newSlice
	default:
		ptr:= reflect.New(reflect.TypeOf(receiver.source))
		elem:=ptr.Elem()
		value:=receiver.val
		numberOfField := value.NumField()

		for i:=0;i<numberOfField;i++{
			elem.Field(i).Set(value.Field(i))
		}

		return reflect.Indirect(elem).Interface()
	}
}