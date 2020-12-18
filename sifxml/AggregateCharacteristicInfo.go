package sifxml

type AggregateCharacteristicInfos []AggregateCharacteristicInfo

type AggregateCharacteristicInfo struct {
	aggregatecharacteristicinfo `xml:"AggregateCharacteristicInfo" json:"AggregateCharacteristicInfo"`
}

type aggregatecharacteristicinfo struct {
	RefId                *RefIdType                `xml:"RefId,attr" json:"RefId"`
	Description          *String                   `xml:"Description,omitempty" json:"Description,omitempty"`
	Definition           *String                   `xml:"Definition" json:"Definition"`
	ElementName          *String                   `xml:"ElementName,omitempty" json:"ElementName,omitempty"`
	LocalCodeList        *LocalCodeListType        `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
	SIF_Metadata         *SIF_MetadataType         `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
	SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
}
