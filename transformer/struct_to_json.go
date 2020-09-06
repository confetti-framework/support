package transformer

import (
	"encoding/json"
	"github.com/lanvard/support"
	"reflect"
)

type StructToJson struct{}

func (j StructToJson) IsValid(object interface{}) bool {
	return object == nil || support.Type(object) == reflect.Struct
}

func (j StructToJson) Transform(object interface{}) string {
	if !j.IsValid(object) {
		panic("can not transform to json with an unsupported type")
	}

	result, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}

	return string(result)
}
