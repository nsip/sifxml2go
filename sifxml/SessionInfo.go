package sifxml

import _ "github.com/creasty/defaults"


type SessionInfos []SessionInfo

    type SessionInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      TimeTableCellRefId string `xml:"TimeTableCellRefId" json:"TimeTableCellRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      LocalId LocalIdType `xml:"LocalId" json:"LocalId"`
      TimeTableSubjectLocalId LocalIdType `xml:"TimeTableSubjectLocalId" json:"TimeTableSubjectLocalId"`
      TeachingGroupLocalId LocalIdType `xml:"TeachingGroupLocalId" json:"TeachingGroupLocalId"`
      SchoolLocalId LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      StaffPersonalLocalId LocalIdType `xml:"StaffPersonalLocalId,omitempty" json:"StaffPersonalLocalId"`
      RoomNumber string `xml:"RoomNumber" json:"RoomNumber"`
      DayId LocalIdType `xml:"DayId" json:"DayId"`
      PeriodId LocalIdType `xml:"PeriodId" json:"PeriodId"`
      SessionDate string `xml:"SessionDate" json:"SessionDate"`
      StartTime string `xml:"StartTime,omitempty" json:"StartTime"`
      FinishTime string `xml:"FinishTime,omitempty" json:"FinishTime"`
      RollMarked AUCodeSetsYesOrNoCategoryType `xml:"RollMarked,omitempty" json:"RollMarked"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    