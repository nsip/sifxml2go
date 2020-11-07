package sifxml

import _ "github.com/creasty/defaults"


type StudentContactPersonals []StudentContactPersonal

    type StudentContactPersonal struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      OtherIdList OtherIdListType `xml:"OtherIdList,omitempty" json:"OtherIdList"`
      PersonInfo PersonInfoType `xml:"PersonInfo" json:"PersonInfo"`
      EmploymentType AUCodeSetsEmploymentTypeType `xml:"EmploymentType,omitempty" json:"EmploymentType"`
      SchoolEducationalLevel EducationalLevelType `xml:"SchoolEducationalLevel,omitempty" json:"SchoolEducationalLevel"`
      NonSchoolEducation AUCodeSetsNonSchoolEducationType `xml:"NonSchoolEducation,omitempty" json:"NonSchoolEducation"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    