package sifxml


type WellbeingAppeals []WellbeingAppeal

    type WellbeingAppeal struct {
  wellbeingappeal `xml:"WellbeingAppeal" json:"WellbeingAppeal"`
}

type wellbeingappeal struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      WellbeingResponseRefId *String `xml:"WellbeingResponseRefId" json:"WellbeingResponseRefId"`
      LocalAppealId *LocalIdType `xml:"LocalAppealId,omitempty" json:"LocalAppealId,omitempty"`
      AppealStatusCode *AUCodeSetsWellbeingAppealStatusType `xml:"AppealStatusCode,omitempty" json:"AppealStatusCode,omitempty"`
      Date *String `xml:"Date,omitempty" json:"Date,omitempty"`
      AppealNotes *String `xml:"AppealNotes,omitempty" json:"AppealNotes,omitempty"`
      AppealOutcome *String `xml:"AppealOutcome,omitempty" json:"AppealOutcome,omitempty"`
      DocumentList *WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

