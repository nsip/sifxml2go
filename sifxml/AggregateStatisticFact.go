package sifxml


type AggregateStatisticFacts []AggregateStatisticFact

    type AggregateStatisticFact struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      AggregateStatisticInfoRefId *String `xml:"AggregateStatisticInfoRefId" json:"AggregateStatisticInfoRefId"`
      Characteristics *CharacteristicsType `xml:"Characteristics" json:"Characteristics"`
      Excluded *String `xml:"Excluded,omitempty" json:"Excluded"`
      Value *Float `xml:"Value" json:"Value"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    