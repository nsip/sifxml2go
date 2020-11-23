package funcapi

//
// public object
//
type CharacteristicsType struct {
	characteristicsType
}

//
// internal object, members exposed only by functions
//
type characteristicsType struct {
	AggregateCharacteristicInfoRefId []string `xml:"AggregateCharacteristicInfoRefId" json:"AggregateCharacteristicInfoRefId"`
}

//
// exposes the internal array
//
func (t *CharacteristicsType) AggregateCharacteristicInfoRefId() *CharacteristicsType {
	if t.characteristicsType.AggregateCharacteristicInfoRefId == nil {
		t.characteristicsType.AggregateCharacteristicInfoRefId = make([]string, 0)
	}
	return t
}

//
// syntactic sugar, not strictly necessary but helps the user
//
func (t *CharacteristicsType) Append() *CharacteristicsType {
	if t.characteristicsType.AggregateCharacteristicInfoRefId == nil {
		t.characteristicsType.AggregateCharacteristicInfoRefId = make([]string, 0)
	}
	return t
}

//
// add a value to the internal array
//
func (t *CharacteristicsType) SetValue(v interface{}) *CharacteristicsType {
	vstr, ok := v.(string)
	if ok {
		t.characteristicsType.AggregateCharacteristicInfoRefId = append(t.characteristicsType.AggregateCharacteristicInfoRefId, vstr)
	}
	return t
}
