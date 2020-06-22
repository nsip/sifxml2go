package sifxml


    type PersonalisedPlan struct {
        RefId RefIdType `xml:"RefId,attr"`
      LocalId LocalIdType `xml:"LocalId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      PersonalisedPlanCategory string `xml:"PersonalisedPlanCategory"`
      PersonalisedPlanStartDate string `xml:"PersonalisedPlanStartDate"`
      PersonalisedPlanEndDate string `xml:"PersonalisedPlanEndDate"`
      PersonalisedPlanReviewDate string `xml:"PersonalisedPlanReviewDate"`
      PersonalisedPlanNotes string `xml:"PersonalisedPlanNotes"`
      DocumentList WellbeingDocumentListType `xml:"DocumentList"`
      AssociatedAttachment string `xml:"AssociatedAttachment"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    