package sifxml


type NAPTestlets []NAPTestlet

    type NAPTestlet struct {
  naptestlet `xml:"NAPTestlet" json:"NAPTestlet"`
}

type naptestlet struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      NAPTestRefId *String `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId,omitempty"`
      NAPTestLocalId *LocalIdType `xml:"NAPTestLocalId" json:"NAPTestLocalId"`
      TestletContent *NAPTestletContentType `xml:"TestletContent" json:"TestletContent"`
      TestItemList *NAPTestItemListType `xml:"TestItemList" json:"TestItemList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

