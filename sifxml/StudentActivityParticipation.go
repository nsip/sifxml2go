package sifxml


type StudentActivityParticipations []StudentActivityParticipation

    type StudentActivityParticipation struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *string `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      StudentActivityInfoRefId *string `xml:"StudentActivityInfoRefId" json:"StudentActivityInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      ParticipationComment *string `xml:"ParticipationComment,omitempty" json:"ParticipationComment"`
      StartDate *string `xml:"StartDate,omitempty" json:"StartDate"`
      EndDate *string `xml:"EndDate,omitempty" json:"EndDate"`
      Role *string `xml:"Role,omitempty" json:"Role"`
      RecognitionList *RecognitionListType `xml:"RecognitionList,omitempty" json:"RecognitionList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    