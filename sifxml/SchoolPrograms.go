package sifxml

import _ "github.com/creasty/defaults"


type SchoolProgramss []SchoolPrograms

    type SchoolPrograms struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      SchoolProgramList SchoolProgramListType `xml:"SchoolProgramList,omitempty" json:"SchoolProgramList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    