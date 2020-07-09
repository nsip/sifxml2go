package sifxml


type WellbeingAlerts []WellbeingAlert

    type WellbeingAlert struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StudentPersonalRefId *string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      Date *string `xml:"Date,omitempty" json:"Date"`
      WellbeingAlertStartDate *string `xml:"WellbeingAlertStartDate,omitempty" json:"WellbeingAlertStartDate"`
      WellbeingAlertEndDate *string `xml:"WellbeingAlertEndDate,omitempty" json:"WellbeingAlertEndDate"`
      WellbeingAlertCategory *string `xml:"WellbeingAlertCategory,omitempty" json:"WellbeingAlertCategory"`
      WellbeingAlertDescription *string `xml:"WellbeingAlertDescription,omitempty" json:"WellbeingAlertDescription"`
      EnrolmentRestricted *string `xml:"EnrolmentRestricted,omitempty" json:"EnrolmentRestricted"`
      AlertAudience *string `xml:"AlertAudience,omitempty" json:"AlertAudience"`
      AlertSeverity *string `xml:"AlertSeverity,omitempty" json:"AlertSeverity"`
      AlertKeyContact *string `xml:"AlertKeyContact,omitempty" json:"AlertKeyContact"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    