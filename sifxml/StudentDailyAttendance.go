package sifxml

import _ "github.com/creasty/defaults"


type StudentDailyAttendances []StudentDailyAttendance

    type StudentDailyAttendance struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      Date string `xml:"Date" json:"Date"`
      SchoolYear SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      DayValue AUCodeSetsDayValueCodeType `xml:"DayValue,omitempty" json:"DayValue"`
      AttendanceCode AttendanceCodeType `xml:"AttendanceCode" json:"AttendanceCode"`
      AttendanceStatus AUCodeSetsAttendanceStatusType `xml:"AttendanceStatus" json:"AttendanceStatus"`
      TimeIn string `xml:"TimeIn" json:"TimeIn"`
      TimeOut string `xml:"TimeOut" json:"TimeOut"`
      AbsenceValue float64 `xml:"AbsenceValue" json:"AbsenceValue"`
      AttendanceNote string `xml:"AttendanceNote,omitempty" json:"AttendanceNote"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    