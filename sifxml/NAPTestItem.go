package sifxml


type NAPTestItems []NAPTestItem

    type NAPTestItem struct {
  naptestitem `xml:"NAPTestItem" json:"NAPTestItem"`
}

type naptestitem struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      TestItemContent *NAPTestItemContentType `xml:"TestItemContent" json:"TestItemContent"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    