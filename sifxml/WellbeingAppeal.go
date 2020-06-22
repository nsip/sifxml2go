package sifxml


    type WellbeingAppeal struct {
        RefId RefIdType `xml:"RefId,attr"`
      LocalId LocalIdType `xml:"LocalId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      WellbeingResponseRefId string `xml:"WellbeingResponseRefId"`
      LocalAppealId LocalIdType `xml:"LocalAppealId"`
      AppealStatusCode string `xml:"AppealStatusCode"`
      Date string `xml:"Date"`
      AppealNotes string `xml:"AppealNotes"`
      AppealOutcome string `xml:"AppealOutcome"`
      DocumentList WellbeingDocumentListType `xml:"DocumentList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    