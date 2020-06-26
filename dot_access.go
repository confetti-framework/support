package support
//
// import "strings"
//
// func Get(target interface{}, key string) Value {
// 	switch target.(type) {
// 	case Value:
// 		return target
// 	}
//
// 	return result, err
// }
//
//
//
// func getUnknown(target interface{}, totalKey string) interface{} {
// 	keys := strings.Split(totalKey, ".")
//
// 	switch target.(type) {
// 	case Value:
// 		return target
// 	case []Value:
// 		return
// 	}
//
// 	return result, err
// }