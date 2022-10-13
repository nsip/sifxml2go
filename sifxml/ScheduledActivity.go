package sifxml


type ScheduledActivitys []ScheduledActivity

    type ScheduledActivity struct {
  scheduledactivity `xml:"ScheduledActivity" json:"ScheduledActivity"`
}

type scheduledactivity struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      TimeTableCellRefId *String `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId,omitempty"`
      DayId *LocalIdType `xml:"DayId,omitempty" json:"DayId,omitempty"`
      PeriodId *LocalIdType `xml:"PeriodId,omitempty" json:"PeriodId,omitempty"`
      TimeTableRefId *String `xml:"TimeTableRefId,omitempty" json:"TimeTableRefId,omitempty"`
      ActivityDate *String `xml:"ActivityDate" json:"ActivityDate"`
      ActivityEndDate *String `xml:"ActivityEndDate,omitempty" json:"ActivityEndDate,omitempty"`
      StartTime *String `xml:"StartTime" json:"StartTime"`
      FinishTime *String `xml:"FinishTime" json:"FinishTime"`
      CellType *String `xml:"CellType,omitempty" json:"CellType,omitempty"`
      TimeTableSubjectRefId *String `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId,omitempty"`
      TeacherList *ScheduledTeacherListType `xml:"TeacherList,omitempty" json:"TeacherList,omitempty"`
      RoomList *RoomListType `xml:"RoomList,omitempty" json:"RoomList,omitempty"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList,omitempty"`
      Location *String `xml:"Location,omitempty" json:"Location,omitempty"`
      ActivityType *AUCodeSetsScheduledActivityTypeType `xml:"ActivityType,omitempty" json:"ActivityType,omitempty"`
      ActivityName *String `xml:"ActivityName,omitempty" json:"ActivityName,omitempty"`
      ActivityComment *String `xml:"ActivityComment,omitempty" json:"ActivityComment,omitempty"`
      StudentList *StudentsType `xml:"StudentList,omitempty" json:"StudentList,omitempty"`
      TeachingGroupList *TeachingGroupListType `xml:"TeachingGroupList,omitempty" json:"TeachingGroupList,omitempty"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels,omitempty"`
      Override *ScheduledActivityOverrideType `xml:"Override,omitempty" json:"Override,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

