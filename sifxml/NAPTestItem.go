package sifxml


type NAPTestItems []NAPTestItem

    type NAPTestItem struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      TestItemContent *NAPTestItemContentType `xml:"TestItemContent,omitempty" json:"TestItemContent"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    