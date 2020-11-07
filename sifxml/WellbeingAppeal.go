package sifxml

import _ "github.com/creasty/defaults"


type WellbeingAppeals []WellbeingAppeal

    type WellbeingAppeal struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      WellbeingResponseRefId string `xml:"WellbeingResponseRefId" json:"WellbeingResponseRefId"`
      LocalAppealId LocalIdType `xml:"LocalAppealId,omitempty" json:"LocalAppealId"`
      AppealStatusCode AUCodeSetsWellbeingAppealStatusType `xml:"AppealStatusCode,omitempty" json:"AppealStatusCode"`
      Date string `xml:"Date,omitempty" json:"Date"`
      AppealNotes string `xml:"AppealNotes,omitempty" json:"AppealNotes"`
      AppealOutcome string `xml:"AppealOutcome,omitempty" json:"AppealOutcome"`
      DocumentList WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    