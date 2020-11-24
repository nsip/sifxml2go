package sifxml


type CalendarSummarys []CalendarSummary

    type CalendarSummary struct {
  calendarsummary `xml:"CalendarSummary" json:"CalendarSummary"`
}

type calendarsummary struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      DaysInSession *Int `xml:"DaysInSession" json:"DaysInSession"`
      StartDate *String `xml:"StartDate,omitempty" json:"StartDate,omitempty"`
      EndDate *String `xml:"EndDate,omitempty" json:"EndDate,omitempty"`
      FirstInstructionDate *String `xml:"FirstInstructionDate,omitempty" json:"FirstInstructionDate,omitempty"`
      LastInstructionDate *String `xml:"LastInstructionDate,omitempty" json:"LastInstructionDate,omitempty"`
      GraduationDate *GraduationDateType `xml:"GraduationDate,omitempty" json:"GraduationDate,omitempty"`
      InstructionalMinutes *Int `xml:"InstructionalMinutes,omitempty" json:"InstructionalMinutes,omitempty"`
      MinutesPerDay *Int `xml:"MinutesPerDay,omitempty" json:"MinutesPerDay,omitempty"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    