package sifxml


    type NAPTestItem struct {
        RefId RefIdType `xml:"RefId,attr"`
      TestItemContent NAPTestItemContentType `xml:"TestItemContent"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    