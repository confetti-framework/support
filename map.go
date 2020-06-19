package support

type Map map[string]Value

func NewMap(itemsRange ...map[string]interface{}) Map {
	result := Map{}

	for _, items := range itemsRange {
		for key, item := range items {
			result[key] = NewValue(item)
		}
	}

	return result
}

// @todo use generics
func NewMapByString(itemsRange ...map[string]string) Map {
	result := Map{}

	for _, items := range itemsRange {
		for key, item := range items {
			result[key] = NewValue(item)
		}
	}

	return result
}
