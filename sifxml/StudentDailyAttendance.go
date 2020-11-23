package sifxml


type StudentDailyAttendances []StudentDailyAttendance

    type StudentDailyAttendance struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      Date *String `xml:"Date" json:"Date"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      DayValue *AUCodeSetsDayValueCodeType `xml:"DayValue,omitempty" json:"DayValue"`
      AttendanceCode *AttendanceCodeType `xml:"AttendanceCode" json:"AttendanceCode"`
      AttendanceStatus *AUCodeSetsAttendanceStatusType `xml:"AttendanceStatus" json:"AttendanceStatus"`
      TimeIn *String `xml:"TimeIn,omitempty" json:"TimeIn"`
      TimeOut *String `xml:"TimeOut,omitempty" json:"TimeOut"`
      AbsenceValue *Float `xml:"AbsenceValue,omitempty" json:"AbsenceValue"`
      AttendanceNote *String `xml:"AttendanceNote,omitempty" json:"AttendanceNote"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    