package sifxml


    type NAPTestlet struct {
        RefId RefIdType `xml:"RefId,attr"`
      NAPTestRefId string `xml:"NAPTestRefId"`
      NAPTestLocalId LocalIdType `xml:"NAPTestLocalId"`
      TestletContent NAPTestletContentType `xml:"TestletContent"`
      TestItemList NAPTestItemListType `xml:"TestItemList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    