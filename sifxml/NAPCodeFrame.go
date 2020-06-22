package sifxml


    type NAPCodeFrame struct {
        RefId RefIdType `xml:"RefId,attr"`
      NAPTestRefId string `xml:"NAPTestRefId"`
      TestContent NAPTestContentType `xml:"TestContent"`
      TestletList NAPCodeFrameTestletListType `xml:"TestletList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    