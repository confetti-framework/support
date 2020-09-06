package transformer

type StructToJson struct {}

func (j StructToJson) CanTransform(object interface{}) bool {
	if object == nil {
		return true
	}

	return false
}

