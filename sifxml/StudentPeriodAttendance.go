package sifxml


type StudentPeriodAttendances []StudentPeriodAttendance

    type StudentPeriodAttendance struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      Date *String `xml:"Date" json:"Date"`
      SessionInfoRefId *String `xml:"SessionInfoRefId,omitempty" json:"SessionInfoRefId"`
      ScheduledActivityRefId *String `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId"`
      TimetablePeriod *String `xml:"TimetablePeriod,omitempty" json:"TimetablePeriod"`
      TimeIn *String `xml:"TimeIn,omitempty" json:"TimeIn"`
      TimeOut *String `xml:"TimeOut,omitempty" json:"TimeOut"`
      AttendanceCode *AttendanceCodeType `xml:"AttendanceCode" json:"AttendanceCode"`
      AttendanceStatus *AUCodeSetsAttendanceStatusType `xml:"AttendanceStatus" json:"AttendanceStatus"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      AuditInfo *AuditInfoType `xml:"AuditInfo,omitempty" json:"AuditInfo"`
      AttendanceComment *String `xml:"AttendanceComment,omitempty" json:"AttendanceComment"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    