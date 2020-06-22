package sifxml


    type WellbeingAlert struct {
        RefId RefIdType `xml:"RefId,attr"`
      LocalId LocalIdType `xml:"LocalId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      Date string `xml:"Date"`
      WellbeingAlertStartDate string `xml:"WellbeingAlertStartDate"`
      WellbeingAlertEndDate string `xml:"WellbeingAlertEndDate"`
      WellbeingAlertCategory string `xml:"WellbeingAlertCategory"`
      WellbeingAlertDescription string `xml:"WellbeingAlertDescription"`
      EnrolmentRestricted string `xml:"EnrolmentRestricted"`
      AlertAudience string `xml:"AlertAudience"`
      AlertSeverity string `xml:"AlertSeverity"`
      AlertKeyContact string `xml:"AlertKeyContact"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    