package sifxml


type TimeTableContainers []TimeTableContainer

    type TimeTableContainer struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      TimeTableSchedule *TimeTableScheduleType `xml:"TimeTableSchedule,omitempty" json:"TimeTableSchedule"`
      TimeTableScheduleCellList *TimeTableScheduleCellListType `xml:"TimeTableScheduleCellList,omitempty" json:"TimeTableScheduleCellList"`
      TeachingGroupScheduleList *TeachingGroupScheduleType `xml:"TeachingGroupScheduleList,omitempty" json:"TeachingGroupScheduleList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    