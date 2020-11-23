package sifxml


type CalendarDates []CalendarDate

    type CalendarDate struct {
        CalendarDateRefId *String `xml:"CalendarDateRefId,attr" json:"CalendarDateRefId"`
      Date *String `xml:"Date" json:"Date"`
      CalendarSummaryRefId *String `xml:"CalendarSummaryRefId,omitempty" json:"CalendarSummaryRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      CalendarDateType *CalendarDateInfoType `xml:"CalendarDateType" json:"CalendarDateType"`
      CalendarDateNumber *Int `xml:"CalendarDateNumber,omitempty" json:"CalendarDateNumber"`
      StudentAttendance *AttendanceInfoType `xml:"StudentAttendance,omitempty" json:"StudentAttendance"`
      TeacherAttendance *AttendanceInfoType `xml:"TeacherAttendance,omitempty" json:"TeacherAttendance"`
      AdministratorAttendance *AttendanceInfoType `xml:"AdministratorAttendance,omitempty" json:"AdministratorAttendance"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    