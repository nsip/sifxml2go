package sifxml

import _ "github.com/creasty/defaults"


type NAPTests []NAPTest

    type NAPTest struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      TestContent NAPTestContentType `xml:"TestContent" json:"TestContent"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    