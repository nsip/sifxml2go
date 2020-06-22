package sifxml


    type StudentSectionEnrollment struct {
        RefId RefIdType `xml:"RefId,attr"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      SectionInfoRefId string `xml:"SectionInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear"`
      EntryDate string `xml:"EntryDate"`
      ExitDate string `xml:"ExitDate"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    