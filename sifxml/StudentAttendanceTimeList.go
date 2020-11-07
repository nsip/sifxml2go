package sifxml


type StudentAttendanceTimeLists []StudentAttendanceTimeList

    type StudentAttendanceTimeList struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentPersonalRefId *string `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      Date *string `xml:"Date" json:"Date"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      AttendanceTimes *AttendanceTimesType `xml:"AttendanceTimes,omitempty" json:"AttendanceTimes"`
      PeriodAttendances *PeriodAttendancesType `xml:"PeriodAttendances,omitempty" json:"PeriodAttendances"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    