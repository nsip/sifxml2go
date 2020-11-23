package funcapi

import (
	"encoding/json"
	"fmt"
)

//
// object visible outside of package, but want all
// properties invisible, accessed only via functions
//
type AggregateStatisticFact struct {
	aggregateStatisticFact `json:"AggregateStatisticsFact"`
}

//
// but need 'exposed' members to marshal/unmarshal into json
// hence this embedded but enexported struct that does all the work
//
type aggregateStatisticFact struct {
	RefId                       *RefIdType                    `xml:"RefId,attr" json:"RefId"`
	AggregateStatisticInfoRefId *string                       `xml:"AggregateStatisticInfoRefId" json:"AggregateStatisticInfoRefId"`
	Characteristics             *CharacteristicsType          `xml:"Characteristics" json:"Characteristics"`
	Excluded                    *string                       `xml:"Excluded,omitempty" json:"Excluded"`
	Value                       *float64                      `xml:"Value" json:"Value"`
	StatsCohortYearLevelList    *StatsCohortYearLevelListType `xml:"StatsCohortYearLevelList,omitempty" json:"StatsCohortYearLevelList"`
}

//
// need 'custom' marshal/unmarshal as data all in embedded object
//
// outer container object also has pesc-compliant single member that
// refelcts the sif object name
//
func (asf *AggregateStatisticFact) UnmarshalJSON(b []byte) error {
	//
	// pescwrapper for outmost-level objects
	//
	type pescWrapper struct{ AggregateStatisticFact aggregateStatisticFact }
	return json.Unmarshal(b, &pescWrapper{AggregateStatisticFact: asf.aggregateStatisticFact})
}

func (asf *AggregateStatisticFact) MarshalJSON() ([]byte, error) {
	//
	// pescwrapper for outmost-level objects
	//
	type pescWrapper struct{ AggregateStatisticFact aggregateStatisticFact }
	return json.Marshal(&pescWrapper{AggregateStatisticFact: asf.aggregateStatisticFact})
}

//
// handle the setting of all properties
// props can be any type and pointers or not
// helper utilities will convert to correct types
// an return pointers as required
//
func (asf *AggregateStatisticFact) SetProperty(key string, value interface{}) *AggregateStatisticFact {

	switch key {
	case "RefId": // need RefIdType pointer
		if v, ok := RefIdPointer(value); ok {
			asf.aggregateStatisticFact.RefId = v
		}
	case "AggregateStatisticInfoRefId": // need string pointer
		if v, ok := StringPointer(value); ok {
			asf.aggregateStatisticFact.AggregateStatisticInfoRefId = v
		}
	case "Excluded": // need string pointer
		if v, ok := StringPointer(value); ok {
			asf.aggregateStatisticFact.Excluded = v
		}
	case "Value": // need float64 pointer
		if v, ok := FloatPointer(value); ok {
			asf.aggregateStatisticFact.Value = v
		}
	default:
		fmt.Printf("Warning: %T has no property: %s\n", asf, key)
	}

	return asf

}

//
// allows for collection of key-value Prop{} objects to be used
// to set many properties at once.
//
func (asf *AggregateStatisticFact) SetProperties(props ...Prop) *AggregateStatisticFact {
	for _, p := range props {
		asf.SetProperty(p.Key, p.Value)
	}
	return asf
}

//
// expose embedded object
//
func (asf *AggregateStatisticFact) Characteristics() *CharacteristicsType {
	if asf.aggregateStatisticFact.Characteristics == nil {
		asf.aggregateStatisticFact.Characteristics = &CharacteristicsType{}
	}
	return asf.aggregateStatisticFact.Characteristics
}

//
// expose embedded object
//
func (asf *AggregateStatisticFact) StatsCohortYearLevelList() *StatsCohortYearLevelListType {
	if asf.aggregateStatisticFact.StatsCohortYearLevelList == nil {
		asf.aggregateStatisticFact.StatsCohortYearLevelList = &StatsCohortYearLevelListType{}
	}
	return asf.aggregateStatisticFact.StatsCohortYearLevelList
}
