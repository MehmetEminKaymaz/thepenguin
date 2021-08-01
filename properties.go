package thepenguin

import "reflect"

func (receiver ReflectionObject) PropertyNames() (result []string) {
	value := receiver.val
	for i:=0;i<value.NumField();i++{
		result = append(result, value.Type().Field(i).Name)
	}
	return result
}
func (receiver ReflectionObject) PropertyValues() (result []interface{}) {
	value := receiver.val
	for i:=0;i<value.NumField();i++{
		result = append(result, value.Field(i).Interface())
	}
	return result
}
func (receiver ReflectionObject) PropertyCount() int  {
	return receiver.val.NumField()
}
func (receiver ReflectionObject) PropertyCountByType(hrefType interface{}) int {
	value := receiver.val
	result:=0
	searchFor := reflect.ValueOf(hrefType).Type()
	for i:=0;i<value.NumField();i++{
		if value.Field(i).Type()==searchFor{
			result++
		}
	}

	return result
}
func (receiver ReflectionObject) PropertyValuesByType(hrefType interface{}) (result []interface{}) {

	value := receiver.val
	searchFor := reflect.ValueOf(hrefType).Type()
	for i:=0;i<value.NumField();i++{
		if value.Field(i).Type()==searchFor{
			result = append(result, value.Field(i).Interface())
		}
	}

	return result
}
func (receiver ReflectionObject) PropertyNamesWithValues() map[string]interface{} {
	value:=receiver.val
	resultMap := make(map[string]interface{})
	for i:=0;i<value.NumField();i++{
		resultMap[value.Type().Field(i).Name]=value.Field(i).Interface()
	}

	return resultMap
}


func (receiver ReflectionObject) PropertyTypeNames() (result []string) {
	value:=receiver.val
	numberOfFields:=value.NumField()
	for i:=0;i<numberOfFields;i++{
		result = append(result, value.Field(i).Type().Name())
	}
	return result
}
func (receiver ReflectionObject) PropertyNamesByKind(kind reflect.Kind)  (result []string) {
	value:=receiver.val
	numberOfFields := value.NumField()
	for i:=0;i<numberOfFields;i++{
		if value.Field(i).Kind()==kind{
			result=append(result,value.Type().Field(i).Name)
		}
	}
	return result
}
func (receiver ReflectionObject) PropertyValuesByKind(kind reflect.Kind)  (result []interface{})  {
	value := receiver.val
	for i:=0;i<value.NumField();i++{
		if value.Field(i).Kind() == kind{
			result = append(result, value.Field(i).Interface())
		}
	}
	return result
}
func (receiver ReflectionObject) PropertyNamesWithValuesByKind( kind reflect.Kind) map[string]interface{} {
	value:=receiver.val
	resultMap := make(map[string]interface{})
	for i:=0;i<value.NumField();i++{
		if value.Field(i).Kind() == kind{
			resultMap[value.Type().Field(i).Name]=value.Field(i).Interface()
		}
	}
	return resultMap
}
