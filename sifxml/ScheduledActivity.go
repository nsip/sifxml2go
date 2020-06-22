package sifxml


    type ScheduledActivity struct {
        RefId RefIdType `xml:"RefId,attr"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      TimeTableCellRefId string `xml:"TimeTableCellRefId"`
      DayId LocalIdType `xml:"DayId"`
      PeriodId LocalIdType `xml:"PeriodId"`
      TimeTableRefId string `xml:"TimeTableRefId"`
      ActivityDate string `xml:"ActivityDate"`
      StartTime string `xml:"StartTime"`
      FinishTime string `xml:"FinishTime"`
      CellType string `xml:"CellType"`
      TimeTableSubjectRefId string `xml:"TimeTableSubjectRefId"`
      TeacherList ScheduledTeacherListType `xml:"TeacherList"`
      RoomList RoomListType `xml:"RoomList"`
      AddressList AddressListType `xml:"AddressList"`
      Location string `xml:"Location"`
      ActivityType string `xml:"ActivityType"`
      ActivityName string `xml:"ActivityName"`
      ActivityComment string `xml:"ActivityComment"`
      StudentList StudentsType `xml:"StudentList"`
      TeachingGroupList TeachingGroupListType `xml:"TeachingGroupList"`
      YearLevels YearLevelsType `xml:"YearLevels"`
      Override ScheduledActivityOverrideType `xml:"Override"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    