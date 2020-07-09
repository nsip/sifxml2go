package sifxml


type TimeTables []TimeTable

    type TimeTable struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      Title *string `xml:"Title,omitempty" json:"Title"`
      DaysPerCycle *int `xml:"DaysPerCycle,omitempty" json:"DaysPerCycle"`
      PeriodsPerDay *int `xml:"PeriodsPerDay,omitempty" json:"PeriodsPerDay"`
      TeachingPeriodsPerDay *int `xml:"TeachingPeriodsPerDay,omitempty" json:"TeachingPeriodsPerDay"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      SchoolName *string `xml:"SchoolName,omitempty" json:"SchoolName"`
      TimeTableCreationDate *string `xml:"TimeTableCreationDate,omitempty" json:"TimeTableCreationDate"`
      StartDate *string `xml:"StartDate,omitempty" json:"StartDate"`
      EndDate *string `xml:"EndDate,omitempty" json:"EndDate"`
      TimeTableDayList *TimeTableDayListType `xml:"TimeTableDayList,omitempty" json:"TimeTableDayList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    