package sifxml


type TimeTables []TimeTable

    type TimeTable struct {
  timetable `xml:"TimeTable" json:"TimeTable"`
}

type timetable struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      Title *String `xml:"Title" json:"Title"`
      DaysPerCycle *Int `xml:"DaysPerCycle" json:"DaysPerCycle"`
      PeriodsPerDay *Int `xml:"PeriodsPerDay" json:"PeriodsPerDay"`
      TeachingPeriodsPerDay *Int `xml:"TeachingPeriodsPerDay,omitempty" json:"TeachingPeriodsPerDay,omitempty"`
      SchoolLocalId *LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId,omitempty"`
      SchoolName *String `xml:"SchoolName,omitempty" json:"SchoolName,omitempty"`
      TimeTableCreationDate *String `xml:"TimeTableCreationDate,omitempty" json:"TimeTableCreationDate,omitempty"`
      StartDate *String `xml:"StartDate,omitempty" json:"StartDate,omitempty"`
      EndDate *String `xml:"EndDate,omitempty" json:"EndDate,omitempty"`
      TimeTableDayList *TimeTableDayListType `xml:"TimeTableDayList" json:"TimeTableDayList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    