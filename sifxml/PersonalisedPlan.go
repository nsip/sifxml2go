package sifxml


type PersonalisedPlans []PersonalisedPlan

    type PersonalisedPlan struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      PersonalisedPlanCategory *AUCodeSetsPersonalisedPlanType `xml:"PersonalisedPlanCategory" json:"PersonalisedPlanCategory"`
      PersonalisedPlanStartDate *String `xml:"PersonalisedPlanStartDate" json:"PersonalisedPlanStartDate"`
      PersonalisedPlanEndDate *String `xml:"PersonalisedPlanEndDate,omitempty" json:"PersonalisedPlanEndDate"`
      PersonalisedPlanReviewDate *String `xml:"PersonalisedPlanReviewDate,omitempty" json:"PersonalisedPlanReviewDate"`
      PersonalisedPlanNotes *String `xml:"PersonalisedPlanNotes,omitempty" json:"PersonalisedPlanNotes"`
      DocumentList *WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList"`
      AssociatedAttachment *String `xml:"AssociatedAttachment,omitempty" json:"AssociatedAttachment"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    