package sifxml


type PersonalisedPlans []PersonalisedPlan

    type PersonalisedPlan struct {
  personalisedplan `xml:"PersonalisedPlan" json:"PersonalisedPlan"`
}

type personalisedplan struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      PersonalisedPlanCategory *AUCodeSetsPersonalisedPlanType `xml:"PersonalisedPlanCategory" json:"PersonalisedPlanCategory"`
      PersonalisedPlanStartDate *String `xml:"PersonalisedPlanStartDate" json:"PersonalisedPlanStartDate"`
      PersonalisedPlanEndDate *String `xml:"PersonalisedPlanEndDate,omitempty" json:"PersonalisedPlanEndDate,omitempty"`
      PersonalisedPlanReviewDate *String `xml:"PersonalisedPlanReviewDate,omitempty" json:"PersonalisedPlanReviewDate,omitempty"`
      PersonalisedPlanNotes *String `xml:"PersonalisedPlanNotes,omitempty" json:"PersonalisedPlanNotes,omitempty"`
      DocumentList *WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList,omitempty"`
      AssociatedAttachment *String `xml:"AssociatedAttachment,omitempty" json:"AssociatedAttachment,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    