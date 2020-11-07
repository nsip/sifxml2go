package sifxml

import _ "github.com/creasty/defaults"


type ScheduledActivitys []ScheduledActivity

    type ScheduledActivity struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      TimeTableCellRefId string `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId"`
      DayId LocalIdType `xml:"DayId,omitempty" json:"DayId"`
      PeriodId LocalIdType `xml:"PeriodId,omitempty" json:"PeriodId"`
      TimeTableRefId string `xml:"TimeTableRefId,omitempty" json:"TimeTableRefId"`
      ActivityDate string `xml:"ActivityDate" json:"ActivityDate"`
      StartTime string `xml:"StartTime" json:"StartTime"`
      FinishTime string `xml:"FinishTime" json:"FinishTime"`
      CellType string `xml:"CellType,omitempty" json:"CellType"`
      TimeTableSubjectRefId string `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      TeacherList ScheduledTeacherListType `xml:"TeacherList,omitempty" json:"TeacherList"`
      RoomList RoomListType `xml:"RoomList,omitempty" json:"RoomList"`
      AddressList AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      Location string `xml:"Location,omitempty" json:"Location"`
      ActivityType AUCodeSetsScheduledActivityTypeType `xml:"ActivityType,omitempty" json:"ActivityType"`
      ActivityName string `xml:"ActivityName,omitempty" json:"ActivityName"`
      ActivityComment string `xml:"ActivityComment,omitempty" json:"ActivityComment"`
      StudentList StudentsType `xml:"StudentList,omitempty" json:"StudentList"`
      TeachingGroupList TeachingGroupListType `xml:"TeachingGroupList,omitempty" json:"TeachingGroupList"`
      YearLevels YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      Override ScheduledActivityOverrideType `xml:"Override,omitempty" json:"Override"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    