package sifxml

import _ "github.com/creasty/defaults"


type StudentPeriodAttendances []StudentPeriodAttendance

    type StudentPeriodAttendance struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      Date string `xml:"Date" json:"Date"`
      SessionInfoRefId string `xml:"SessionInfoRefId,omitempty" json:"SessionInfoRefId"`
      ScheduledActivityRefId string `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId"`
      TimetablePeriod string `xml:"TimetablePeriod,omitempty" json:"TimetablePeriod"`
      TimeIn string `xml:"TimeIn,omitempty" json:"TimeIn"`
      TimeOut string `xml:"TimeOut,omitempty" json:"TimeOut"`
      AttendanceCode AttendanceCodeType `xml:"AttendanceCode" json:"AttendanceCode"`
      AttendanceStatus AUCodeSetsAttendanceStatusType `xml:"AttendanceStatus" json:"AttendanceStatus"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      AuditInfo AuditInfoType `xml:"AuditInfo,omitempty" json:"AuditInfo"`
      AttendanceComment string `xml:"AttendanceComment,omitempty" json:"AttendanceComment"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    