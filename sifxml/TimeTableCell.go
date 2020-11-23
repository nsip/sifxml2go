package sifxml


type TimeTableCells []TimeTableCell

    type TimeTableCell struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      TimeTableRefId *String `xml:"TimeTableRefId" json:"TimeTableRefId"`
      TimeTableSubjectRefId *String `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      TeachingGroupRefId *String `xml:"TeachingGroupRefId" json:"TeachingGroupRefId"`
      RoomInfoRefId *String `xml:"RoomInfoRefId,omitempty" json:"RoomInfoRefId"`
      StaffPersonalRefId *String `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      TimeTableLocalId *LocalIdType `xml:"TimeTableLocalId,omitempty" json:"TimeTableLocalId"`
      SubjectLocalId *LocalIdType `xml:"SubjectLocalId,omitempty" json:"SubjectLocalId"`
      TeachingGroupLocalId *LocalIdType `xml:"TeachingGroupLocalId,omitempty" json:"TeachingGroupLocalId"`
      RoomNumber *HomeroomNumberType `xml:"RoomNumber,omitempty" json:"RoomNumber"`
      StaffLocalId *LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId"`
      DayId *LocalIdType `xml:"DayId" json:"DayId"`
      PeriodId *LocalIdType `xml:"PeriodId" json:"PeriodId"`
      CellType *String `xml:"CellType" json:"CellType"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      TeacherList *ScheduledTeacherListType `xml:"TeacherList,omitempty" json:"TeacherList"`
      RoomList *RoomListType `xml:"RoomList,omitempty" json:"RoomList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    