package funcapi

import (
	"fmt"
)

//
// externally visible object
//
type YearLevelType struct {
	yearLevelType
}

//
// internal object
//
type yearLevelType struct {
	Code *AUCodeSetsYearLevelCodeType `xml:"Code" json:"Code"`
}

//
// allows user to set internal property, in this case shows how
// codesets can be checked
//
func (y *YearLevelType) SetProperty(key string, value interface{}) *YearLevelType {

	switch key {
	case "Code":
		if v, ok := AUCodeSetsYearLevelCodeTypePointer(value); ok {
			if CodesetContains(AUCodeSetsYearLevelCodeTypeValues, value) {
				y.Code = v
			} else {
				fmt.Printf("Warning: %v is not a valid value for %T.Code\n", value, y)
			}
		}
	default:
		fmt.Printf("Warning: %T has no property: %s\n", y, key)
	}

	return y
}
