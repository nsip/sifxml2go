package sifxml


type SessionInfos []SessionInfo

    type SessionInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      TimeTableCellRefId *String `xml:"TimeTableCellRefId" json:"TimeTableCellRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      TimeTableSubjectLocalId *LocalIdType `xml:"TimeTableSubjectLocalId" json:"TimeTableSubjectLocalId"`
      TeachingGroupLocalId *LocalIdType `xml:"TeachingGroupLocalId" json:"TeachingGroupLocalId"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      StaffPersonalLocalId *LocalIdType `xml:"StaffPersonalLocalId,omitempty" json:"StaffPersonalLocalId"`
      RoomNumber *String `xml:"RoomNumber,omitempty" json:"RoomNumber"`
      DayId *LocalIdType `xml:"DayId" json:"DayId"`
      PeriodId *LocalIdType `xml:"PeriodId" json:"PeriodId"`
      SessionDate *String `xml:"SessionDate" json:"SessionDate"`
      StartTime *String `xml:"StartTime,omitempty" json:"StartTime"`
      FinishTime *String `xml:"FinishTime,omitempty" json:"FinishTime"`
      RollMarked *AUCodeSetsYesOrNoCategoryType `xml:"RollMarked,omitempty" json:"RollMarked"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    