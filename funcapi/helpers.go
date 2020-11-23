package funcapi

import "fmt"

//
// indirection types from original generated code
//
type GUIDType string
type MsgIdType GUIDType
type RefIdType GUIDType
type LocalIdType string
type AUCodeSetsYearLevelCodeType string

//
// wrapper for key/val property
//
type Prop struct {
	Key   string
	Value interface{}
}

//
// series of XXPointer methods to coerce interfaces to
// required types and to return pointer as required
//

//
//
func StringPointer(value interface{}) (*string, bool) {

	switch t := value.(type) {
	case *string:
		return value.(*string), true
	case string:
		v, _ := value.(string)
		return &v, true
	default:
		fmt.Printf("Warning: cannot resolve %T (%v) to string\n", t, value)
	}

	return nil, false
}

//
// allows strings as well as derived types
//
func RefIdPointer(value interface{}) (*RefIdType, bool) {

	switch t := value.(type) {
	case *RefIdType:
		return value.(*RefIdType), true
	case RefIdType:
		v, _ := value.(RefIdType)
		return &v, true
	case string:
		vstr, _ := value.(string)
		v := RefIdType(vstr)
		return &v, true
	default:
		fmt.Printf("Warning: cannot resolve %T (%v) to RefId\n", t, value)
	}

	return nil, false
}

//
// allows strings as well as derived types
//
func LocalIdPointer(value interface{}) (*LocalIdType, bool) {

	switch t := value.(type) {
	case *LocalIdType:
		return value.(*LocalIdType), true
	case LocalIdType:
		v, _ := value.(LocalIdType)
		return &v, true
	case string:
		vstr, _ := value.(string)
		v := LocalIdType(vstr)
		return &v, true
	default:
		fmt.Printf("Warning: cannot resolve %T (%v) to LocalId\n", t, value)
	}

	return nil, false
}

//
//
func IntPointer(value interface{}) (*int, bool) {

	switch t := value.(type) {
	case *int:
		return value.(*int), true
	case int:
		v, _ := value.(int)
		return &v, true
	default:
		fmt.Printf("Warning: cannot resolve %T (%v) to int\n", t, value)
	}

	return nil, false
}

//
// allows float32 as well as 64
//
func FloatPointer(value interface{}) (*float64, bool) {

	switch t := value.(type) {
	case *float64:
		return value.(*float64), true
	case float64:
		v, _ := value.(float64)
		return &v, true
	case float32:
		v32, _ := value.(float32)
		v64 := float64(v32)
		return &v64, true
	default:
		fmt.Printf("Warning: cannot resolve %T (%v) to float64\n", t, value)
	}

	return nil, false
}

//
//
func AUCodeSetsYearLevelCodeTypePointer(value interface{}) (*AUCodeSetsYearLevelCodeType, bool) {

	switch t := value.(type) {
	case *AUCodeSetsYearLevelCodeType:
		return value.(*AUCodeSetsYearLevelCodeType), true
	case AUCodeSetsYearLevelCodeType:
		v, _ := value.(AUCodeSetsYearLevelCodeType)
		return &v, true
	case string:
		vstr, _ := value.(string)
		v := AUCodeSetsYearLevelCodeType(vstr)
		return &v, true
	default:
		fmt.Printf("Warning: cannot resolve %T (%v) to AUCodeSetsYearLevelCodeType\n", t, value)
	}

	return nil, false
}
