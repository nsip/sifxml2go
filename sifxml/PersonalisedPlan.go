package sifxml


type PersonalisedPlans []PersonalisedPlan

    type PersonalisedPlan struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StudentPersonalRefId *string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      PersonalisedPlanCategory *string `xml:"PersonalisedPlanCategory,omitempty" json:"PersonalisedPlanCategory"`
      PersonalisedPlanStartDate *string `xml:"PersonalisedPlanStartDate,omitempty" json:"PersonalisedPlanStartDate"`
      PersonalisedPlanEndDate *string `xml:"PersonalisedPlanEndDate,omitempty" json:"PersonalisedPlanEndDate"`
      PersonalisedPlanReviewDate *string `xml:"PersonalisedPlanReviewDate,omitempty" json:"PersonalisedPlanReviewDate"`
      PersonalisedPlanNotes *string `xml:"PersonalisedPlanNotes,omitempty" json:"PersonalisedPlanNotes"`
      DocumentList *WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList"`
      AssociatedAttachment *string `xml:"AssociatedAttachment,omitempty" json:"AssociatedAttachment"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    