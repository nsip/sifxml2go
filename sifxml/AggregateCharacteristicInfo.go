package sifxml


type AggregateCharacteristicInfos []AggregateCharacteristicInfo

    type AggregateCharacteristicInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      Description *String `xml:"Description,omitempty" json:"Description"`
      Definition *String `xml:"Definition" json:"Definition"`
      ElementName *String `xml:"ElementName,omitempty" json:"ElementName"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    