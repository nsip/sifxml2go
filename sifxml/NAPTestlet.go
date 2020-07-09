package sifxml


type NAPTestlets []NAPTestlet

    type NAPTestlet struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      NAPTestRefId *string `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId"`
      NAPTestLocalId *LocalIdType `xml:"NAPTestLocalId,omitempty" json:"NAPTestLocalId"`
      TestletContent *NAPTestletContentType `xml:"TestletContent,omitempty" json:"TestletContent"`
      TestItemList *NAPTestItemListType `xml:"TestItemList,omitempty" json:"TestItemList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    