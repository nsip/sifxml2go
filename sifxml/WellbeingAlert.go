package sifxml


type WellbeingAlerts []WellbeingAlert

    type WellbeingAlert struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      Date *String `xml:"Date" json:"Date"`
      WellbeingAlertStartDate *String `xml:"WellbeingAlertStartDate,omitempty" json:"WellbeingAlertStartDate"`
      WellbeingAlertEndDate *String `xml:"WellbeingAlertEndDate,omitempty" json:"WellbeingAlertEndDate"`
      WellbeingAlertCategory *AUCodeSetsWellbeingAlertCategoryType `xml:"WellbeingAlertCategory,omitempty" json:"WellbeingAlertCategory"`
      WellbeingAlertDescription *String `xml:"WellbeingAlertDescription,omitempty" json:"WellbeingAlertDescription"`
      EnrolmentRestricted *AUCodeSetsYesOrNoCategoryType `xml:"EnrolmentRestricted,omitempty" json:"EnrolmentRestricted"`
      AlertAudience *String `xml:"AlertAudience,omitempty" json:"AlertAudience"`
      AlertSeverity *String `xml:"AlertSeverity,omitempty" json:"AlertSeverity"`
      AlertKeyContact *String `xml:"AlertKeyContact,omitempty" json:"AlertKeyContact"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    