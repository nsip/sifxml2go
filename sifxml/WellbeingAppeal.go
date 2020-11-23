package sifxml


type WellbeingAppeals []WellbeingAppeal

    type WellbeingAppeal struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      WellbeingResponseRefId *String `xml:"WellbeingResponseRefId" json:"WellbeingResponseRefId"`
      LocalAppealId *LocalIdType `xml:"LocalAppealId,omitempty" json:"LocalAppealId"`
      AppealStatusCode *AUCodeSetsWellbeingAppealStatusType `xml:"AppealStatusCode,omitempty" json:"AppealStatusCode"`
      Date *String `xml:"Date,omitempty" json:"Date"`
      AppealNotes *String `xml:"AppealNotes,omitempty" json:"AppealNotes"`
      AppealOutcome *String `xml:"AppealOutcome,omitempty" json:"AppealOutcome"`
      DocumentList *WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    