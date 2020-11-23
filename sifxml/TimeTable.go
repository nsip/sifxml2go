package sifxml


type TimeTables []TimeTable

    type TimeTable struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      Title *String `xml:"Title" json:"Title"`
      DaysPerCycle *Int `xml:"DaysPerCycle" json:"DaysPerCycle"`
      PeriodsPerDay *Int `xml:"PeriodsPerDay" json:"PeriodsPerDay"`
      TeachingPeriodsPerDay *Int `xml:"TeachingPeriodsPerDay,omitempty" json:"TeachingPeriodsPerDay"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      SchoolName *String `xml:"SchoolName,omitempty" json:"SchoolName"`
      TimeTableCreationDate *String `xml:"TimeTableCreationDate,omitempty" json:"TimeTableCreationDate"`
      StartDate *String `xml:"StartDate,omitempty" json:"StartDate"`
      EndDate *String `xml:"EndDate,omitempty" json:"EndDate"`
      TimeTableDayList *TimeTableDayListType `xml:"TimeTableDayList" json:"TimeTableDayList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    