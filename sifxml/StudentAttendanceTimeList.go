package sifxml


type StudentAttendanceTimeLists []StudentAttendanceTimeList

    type StudentAttendanceTimeList struct {
  studentattendancetimelist `xml:"StudentAttendanceTimeList" json:"StudentAttendanceTimeList"`
}

type studentattendancetimelist struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      Date *String `xml:"Date" json:"Date"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      AttendanceTimes *AttendanceTimesType `xml:"AttendanceTimes,omitempty" json:"AttendanceTimes,omitempty"`
      PeriodAttendances *PeriodAttendancesType `xml:"PeriodAttendances,omitempty" json:"PeriodAttendances,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    