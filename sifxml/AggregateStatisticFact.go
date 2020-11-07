package sifxml

import _ "github.com/creasty/defaults"


type AggregateStatisticFacts []AggregateStatisticFact

    type AggregateStatisticFact struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      AggregateStatisticInfoRefId string `xml:"AggregateStatisticInfoRefId" json:"AggregateStatisticInfoRefId"`
      Characteristics CharacteristicsType `xml:"Characteristics" json:"Characteristics"`
      Excluded string `xml:"Excluded,omitempty" json:"Excluded"`
      Value float64 `xml:"Value" json:"Value"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    