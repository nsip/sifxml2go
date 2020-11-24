package sifxml


type StudentActivityInfos []StudentActivityInfo

    type StudentActivityInfo struct {
  studentactivityinfo `xml:"StudentActivityInfo" json:"StudentActivityInfo"`
}

type studentactivityinfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      Title *String `xml:"Title" json:"Title"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      StudentActivityType *StudentActivityType `xml:"StudentActivityType" json:"StudentActivityType"`
      StudentActivityLevel *String `xml:"StudentActivityLevel,omitempty" json:"StudentActivityLevel,omitempty"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels,omitempty"`
      CurricularStatus *AUCodeSetsActivityTypeType `xml:"CurricularStatus,omitempty" json:"CurricularStatus,omitempty"`
      Location *LocationType `xml:"Location,omitempty" json:"Location,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    