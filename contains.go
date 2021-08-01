package thepenguin

func (receiver ReflectionObject) Contains(item interface{}) bool {
	value := receiver.val
	length := value.Len()
	for i:=0;i<length;i++{
		if value.Index(i).Interface()==item{
			return true
		}
	}

	return false
}