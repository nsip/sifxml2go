package sifxml


    type TimeTableContainer struct {
        RefId RefIdType `xml:"RefId,attr"`
      TimeTableSchedule TimeTableScheduleType `xml:"TimeTableSchedule"`
      TimeTableScheduleCellList TimeTableScheduleCellListType `xml:"TimeTableScheduleCellList"`
      TeachingGroupScheduleList TeachingGroupScheduleType `xml:"TeachingGroupScheduleList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    