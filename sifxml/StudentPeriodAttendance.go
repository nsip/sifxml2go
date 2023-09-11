package sifxml



    type StudentPeriodAttendances struct {
      studentperiodattendances `json:"StudentPeriodAttendances"`
    }

    type studentperiodattendances struct {
      StudentPeriodAttendance []studentperiodattendance `json:"StudentPeriodAttendance"`
    }

    type StudentPeriodAttendance struct {
      studentperiodattendance `xml:"StudentPeriodAttendance" json:"StudentPeriodAttendance"`
     }

     type studentperiodattendance struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      Date *String `xml:"Date" json:"Date"`
      SessionInfoRefId *String `xml:"SessionInfoRefId,omitempty" json:"SessionInfoRefId,omitempty"`
      ScheduledActivityRefId *String `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId,omitempty"`
      TimetablePeriod *String `xml:"TimetablePeriod,omitempty" json:"TimetablePeriod,omitempty"`
      TimeIn *String `xml:"TimeIn,omitempty" json:"TimeIn,omitempty"`
      TimeOut *String `xml:"TimeOut,omitempty" json:"TimeOut,omitempty"`
      AttendanceCode *AttendanceCodeType `xml:"AttendanceCode" json:"AttendanceCode"`
      AttendanceStatus *AUCodeSetsAttendanceStatusType `xml:"AttendanceStatus" json:"AttendanceStatus"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear,omitempty"`
      AuditInfo *AuditInfoType `xml:"AuditInfo,omitempty" json:"AuditInfo,omitempty"`
      AttendanceComment *String `xml:"AttendanceComment,omitempty" json:"AttendanceComment,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

