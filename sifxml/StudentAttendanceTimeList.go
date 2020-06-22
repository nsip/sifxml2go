package sifxml


    type StudentAttendanceTimeList struct {
        RefId RefIdType `xml:"RefId,attr"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      Date string `xml:"Date"`
      SchoolYear SchoolYearType `xml:"SchoolYear"`
      AttendanceTimes AttendanceTimesType `xml:"AttendanceTimes"`
      PeriodAttendances PeriodAttendancesType `xml:"PeriodAttendances"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    