package sifxml


type ScheduledActivitys []ScheduledActivity

    type ScheduledActivity struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      TimeTableCellRefId *String `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId"`
      DayId *LocalIdType `xml:"DayId,omitempty" json:"DayId"`
      PeriodId *LocalIdType `xml:"PeriodId,omitempty" json:"PeriodId"`
      TimeTableRefId *String `xml:"TimeTableRefId,omitempty" json:"TimeTableRefId"`
      ActivityDate *String `xml:"ActivityDate" json:"ActivityDate"`
      StartTime *String `xml:"StartTime" json:"StartTime"`
      FinishTime *String `xml:"FinishTime" json:"FinishTime"`
      CellType *String `xml:"CellType,omitempty" json:"CellType"`
      TimeTableSubjectRefId *String `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      TeacherList *ScheduledTeacherListType `xml:"TeacherList,omitempty" json:"TeacherList"`
      RoomList *RoomListType `xml:"RoomList,omitempty" json:"RoomList"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      Location *String `xml:"Location,omitempty" json:"Location"`
      ActivityType *AUCodeSetsScheduledActivityTypeType `xml:"ActivityType,omitempty" json:"ActivityType"`
      ActivityName *String `xml:"ActivityName,omitempty" json:"ActivityName"`
      ActivityComment *String `xml:"ActivityComment,omitempty" json:"ActivityComment"`
      StudentList *StudentsType `xml:"StudentList,omitempty" json:"StudentList"`
      TeachingGroupList *TeachingGroupListType `xml:"TeachingGroupList,omitempty" json:"TeachingGroupList"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      Override *ScheduledActivityOverrideType `xml:"Override,omitempty" json:"Override"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    