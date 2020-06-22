package sifxml


    type TimeTableCell struct {
        RefId RefIdType `xml:"RefId,attr"`
      TimeTableRefId string `xml:"TimeTableRefId"`
      TimeTableSubjectRefId string `xml:"TimeTableSubjectRefId"`
      TeachingGroupRefId string `xml:"TeachingGroupRefId"`
      RoomInfoRefId string `xml:"RoomInfoRefId"`
      StaffPersonalRefId string `xml:"StaffPersonalRefId"`
      TimeTableLocalId LocalIdType `xml:"TimeTableLocalId"`
      SubjectLocalId LocalIdType `xml:"SubjectLocalId"`
      TeachingGroupLocalId LocalIdType `xml:"TeachingGroupLocalId"`
      RoomNumber HomeroomNumberType `xml:"RoomNumber"`
      StaffLocalId LocalIdType `xml:"StaffLocalId"`
      DayId LocalIdType `xml:"DayId"`
      PeriodId LocalIdType `xml:"PeriodId"`
      CellType string `xml:"CellType"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      SchoolLocalId LocalIdType `xml:"SchoolLocalId"`
      TeacherList ScheduledTeacherListType `xml:"TeacherList"`
      RoomList RoomListType `xml:"RoomList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    