package sifxml



    type NAPTests struct {
      naptests `json:"NAPTests"`
    }

    type naptests struct {
      NAPTest []naptest `json:"NAPTest"`
    }

    type NAPTest struct {
      naptest `xml:"NAPTest" json:"NAPTest"`
     }

     type naptest struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      TestContent *NAPTestContentType `xml:"TestContent" json:"TestContent"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

