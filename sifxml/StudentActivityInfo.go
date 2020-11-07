package sifxml


type StudentActivityInfos []StudentActivityInfo

    type StudentActivityInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      Title *string `xml:"Title" json:"Title"`
      Description *string `xml:"Description,omitempty" json:"Description"`
      StudentActivityType *StudentActivityType `xml:"StudentActivityType" json:"StudentActivityType"`
      StudentActivityLevel *string `xml:"StudentActivityLevel,omitempty" json:"StudentActivityLevel"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      CurricularStatus *AUCodeSetsActivityTypeType `xml:"CurricularStatus,omitempty" json:"CurricularStatus"`
      Location *LocationType `xml:"Location,omitempty" json:"Location"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    