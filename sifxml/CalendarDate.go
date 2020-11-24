package sifxml


type CalendarDates []CalendarDate

    type CalendarDate struct {
  calendardate `xml:"CalendarDate" json:"CalendarDate"`
}

type calendardate struct {
        CalendarDateRefId *String `xml:"CalendarDateRefId,attr" json:"CalendarDateRefId"`
      Date *String `xml:"Date" json:"Date"`
      CalendarSummaryRefId *String `xml:"CalendarSummaryRefId,omitempty" json:"CalendarSummaryRefId,omitempty"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      CalendarDateType *CalendarDateInfoType `xml:"CalendarDateType" json:"CalendarDateType"`
      CalendarDateNumber *Int `xml:"CalendarDateNumber,omitempty" json:"CalendarDateNumber,omitempty"`
      StudentAttendance *AttendanceInfoType `xml:"StudentAttendance,omitempty" json:"StudentAttendance,omitempty"`
      TeacherAttendance *AttendanceInfoType `xml:"TeacherAttendance,omitempty" json:"TeacherAttendance,omitempty"`
      AdministratorAttendance *AttendanceInfoType `xml:"AdministratorAttendance,omitempty" json:"AdministratorAttendance,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    