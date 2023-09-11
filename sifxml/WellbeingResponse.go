package sifxml



    type WellbeingResponses struct {
      wellbeingresponses `json:"WellbeingResponses"`
    }

    type wellbeingresponses struct {
      WellbeingResponse []wellbeingresponse `json:"WellbeingResponse"`
    }

    type WellbeingResponse struct {
      wellbeingresponse `xml:"WellbeingResponse" json:"WellbeingResponse"`
     }

     type wellbeingresponse struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      Date *String `xml:"Date" json:"Date"`
      WellbeingResponseStartDate *String `xml:"WellbeingResponseStartDate" json:"WellbeingResponseStartDate"`
      WellbeingResponseEndDate *String `xml:"WellbeingResponseEndDate,omitempty" json:"WellbeingResponseEndDate,omitempty"`
      WellbeingResponseCategory *AUCodeSetsWellbeingResponseCategoryType `xml:"WellbeingResponseCategory,omitempty" json:"WellbeingResponseCategory,omitempty"`
      WellbeingResponseNotes *String `xml:"WellbeingResponseNotes,omitempty" json:"WellbeingResponseNotes,omitempty"`
      PersonInvolvementList *PersonInvolvementListType `xml:"PersonInvolvementList,omitempty" json:"PersonInvolvementList,omitempty"`
      Suspension *SuspensionContainerType `xml:"Suspension,omitempty" json:"Suspension,omitempty"`
      Detention *DetentionContainerType `xml:"Detention,omitempty" json:"Detention,omitempty"`
      PlanRequired *PlanRequiredContainerType `xml:"PlanRequired,omitempty" json:"PlanRequired,omitempty"`
      Award *AwardContainerType `xml:"Award,omitempty" json:"Award,omitempty"`
      OtherResponse *OtherWellbeingResponseContainerType `xml:"OtherResponse,omitempty" json:"OtherResponse,omitempty"`
      DocumentList *WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

