package sifxml



    type TimeTableContainers struct {
      timetablecontainers `json:"TimeTableContainers"`
    }

    type timetablecontainers struct {
      TimeTableContainer []timetablecontainer `json:"TimeTableContainer"`
    }

    type TimeTableContainer struct {
      timetablecontainer `xml:"TimeTableContainer" json:"TimeTableContainer"`
     }

     type timetablecontainer struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      TimeTableSchedule *TimeTableScheduleType `xml:"TimeTableSchedule,omitempty" json:"TimeTableSchedule,omitempty"`
      TimeTableScheduleCellList *TimeTableScheduleCellListType `xml:"TimeTableScheduleCellList,omitempty" json:"TimeTableScheduleCellList,omitempty"`
      TeachingGroupScheduleList *TeachingGroupScheduleType `xml:"TeachingGroupScheduleList,omitempty" json:"TeachingGroupScheduleList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

