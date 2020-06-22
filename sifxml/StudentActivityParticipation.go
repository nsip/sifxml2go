package sifxml


    type StudentActivityParticipation struct {
        RefId RefIdType `xml:"RefId,attr"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      StudentActivityInfoRefId string `xml:"StudentActivityInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear"`
      ParticipationComment string `xml:"ParticipationComment"`
      StartDate string `xml:"StartDate"`
      EndDate string `xml:"EndDate"`
      Role string `xml:"Role"`
      RecognitionList RecognitionListType `xml:"RecognitionList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    