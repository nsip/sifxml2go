package sifxml


    type StudentPeriodAttendance struct {
        RefId RefIdType `xml:"RefId,attr"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      Date string `xml:"Date"`
      SessionInfoRefId string `xml:"SessionInfoRefId"`
      ScheduledActivityRefId string `xml:"ScheduledActivityRefId"`
      TimetablePeriod string `xml:"TimetablePeriod"`
      TimeIn string `xml:"TimeIn"`
      TimeOut string `xml:"TimeOut"`
      AttendanceCode AttendanceCodeType `xml:"AttendanceCode"`
      AttendanceStatus string `xml:"AttendanceStatus"`
      SchoolYear SchoolYearType `xml:"SchoolYear"`
      AuditInfo AuditInfoType `xml:"AuditInfo"`
      AttendanceComment string `xml:"AttendanceComment"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    