package sifxml


    type CalendarDate struct {
        CalendarDateRefId string `xml:"CalendarDateRefId,attr"`
      Date string `xml:"Date"`
      CalendarSummaryRefId string `xml:"CalendarSummaryRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear"`
      CalendarDateType CalendarDateInfoType `xml:"CalendarDateType"`
      CalendarDateNumber string `xml:"CalendarDateNumber"`
      StudentAttendance AttendanceInfoType `xml:"StudentAttendance"`
      TeacherAttendance AttendanceInfoType `xml:"TeacherAttendance"`
      AdministratorAttendance AttendanceInfoType `xml:"AdministratorAttendance"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    