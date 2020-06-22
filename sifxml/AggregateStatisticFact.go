package sifxml


    type AggregateStatisticFact struct {
        RefId RefIdType `xml:"RefId,attr"`
      AggregateStatisticInfoRefId string `xml:"AggregateStatisticInfoRefId"`
      Characteristics CharacteristicsType `xml:"Characteristics"`
      Excluded string `xml:"Excluded"`
      Value string `xml:"Value"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    