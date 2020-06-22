package sifxml


    type NAPTest struct {
        RefId RefIdType `xml:"RefId,attr"`
      TestContent NAPTestContentType `xml:"TestContent"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    