package sifxml


type StudentActivityParticipations []StudentActivityParticipation

    type StudentActivityParticipation struct {
  studentactivityparticipation `xml:"StudentActivityParticipation" json:"StudentActivityParticipation"`
}

type studentactivityparticipation struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      StudentActivityInfoRefId *String `xml:"StudentActivityInfoRefId" json:"StudentActivityInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear,omitempty"`
      ParticipationComment *String `xml:"ParticipationComment,omitempty" json:"ParticipationComment,omitempty"`
      StartDate *String `xml:"StartDate,omitempty" json:"StartDate,omitempty"`
      EndDate *String `xml:"EndDate,omitempty" json:"EndDate,omitempty"`
      Role *String `xml:"Role,omitempty" json:"Role,omitempty"`
      RecognitionList *RecognitionListType `xml:"RecognitionList,omitempty" json:"RecognitionList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

