package sifxml


type WellbeingResponses []WellbeingResponse

    type WellbeingResponse struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StudentPersonalRefId *string `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      Date *string `xml:"Date" json:"Date"`
      WellbeingResponseStartDate *string `xml:"WellbeingResponseStartDate" json:"WellbeingResponseStartDate"`
      WellbeingResponseEndDate *string `xml:"WellbeingResponseEndDate,omitempty" json:"WellbeingResponseEndDate"`
      WellbeingResponseCategory *AUCodeSetsWellbeingResponseCategoryType `xml:"WellbeingResponseCategory" json:"WellbeingResponseCategory"`
      WellbeingResponseNotes *string `xml:"WellbeingResponseNotes,omitempty" json:"WellbeingResponseNotes"`
      PersonInvolvementList *PersonInvolvementListType `xml:"PersonInvolvementList,omitempty" json:"PersonInvolvementList"`
      Suspension *SuspensionContainerType `xml:"Suspension,omitempty" json:"Suspension"`
      Detention *DetentionContainerType `xml:"Detention,omitempty" json:"Detention"`
      PlanRequired *PlanRequiredContainerType `xml:"PlanRequired,omitempty" json:"PlanRequired"`
      Award *AwardContainerType `xml:"Award,omitempty" json:"Award"`
      OtherResponse *OtherWellbeingResponseContainerType `xml:"OtherResponse,omitempty" json:"OtherResponse"`
      DocumentList *WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    