package sifxml


type WellbeingAlerts []WellbeingAlert

    type WellbeingAlert struct {
  wellbeingalert `xml:"WellbeingAlert" json:"WellbeingAlert"`
}

type wellbeingalert struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      Date *String `xml:"Date" json:"Date"`
      WellbeingAlertStartDate *String `xml:"WellbeingAlertStartDate,omitempty" json:"WellbeingAlertStartDate,omitempty"`
      WellbeingAlertEndDate *String `xml:"WellbeingAlertEndDate,omitempty" json:"WellbeingAlertEndDate,omitempty"`
      WellbeingAlertCategory *AUCodeSetsWellbeingAlertCategoryType `xml:"WellbeingAlertCategory,omitempty" json:"WellbeingAlertCategory,omitempty"`
      WellbeingAlertDescription *String `xml:"WellbeingAlertDescription,omitempty" json:"WellbeingAlertDescription,omitempty"`
      EnrolmentRestricted *AUCodeSetsYesOrNoCategoryType `xml:"EnrolmentRestricted,omitempty" json:"EnrolmentRestricted,omitempty"`
      AlertAudience *String `xml:"AlertAudience,omitempty" json:"AlertAudience,omitempty"`
      AlertSeverity *String `xml:"AlertSeverity,omitempty" json:"AlertSeverity,omitempty"`
      AlertKeyContact *String `xml:"AlertKeyContact,omitempty" json:"AlertKeyContact,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

