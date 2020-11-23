package sifxml


type CalendarSummarys []CalendarSummary

    type CalendarSummary struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      Description *String `xml:"Description,omitempty" json:"Description"`
      DaysInSession *Int `xml:"DaysInSession" json:"DaysInSession"`
      StartDate *String `xml:"StartDate,omitempty" json:"StartDate"`
      EndDate *String `xml:"EndDate,omitempty" json:"EndDate"`
      FirstInstructionDate *String `xml:"FirstInstructionDate,omitempty" json:"FirstInstructionDate"`
      LastInstructionDate *String `xml:"LastInstructionDate,omitempty" json:"LastInstructionDate"`
      GraduationDate *GraduationDateType `xml:"GraduationDate,omitempty" json:"GraduationDate"`
      InstructionalMinutes *Int `xml:"InstructionalMinutes,omitempty" json:"InstructionalMinutes"`
      MinutesPerDay *Int `xml:"MinutesPerDay,omitempty" json:"MinutesPerDay"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    