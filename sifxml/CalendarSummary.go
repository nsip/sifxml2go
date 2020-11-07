package sifxml


type CalendarSummarys []CalendarSummary

    type CalendarSummary struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      Description *string `xml:"Description,omitempty" json:"Description"`
      DaysInSession *int `xml:"DaysInSession" json:"DaysInSession"`
      StartDate *string `xml:"StartDate,omitempty" json:"StartDate"`
      EndDate *string `xml:"EndDate,omitempty" json:"EndDate"`
      FirstInstructionDate *string `xml:"FirstInstructionDate,omitempty" json:"FirstInstructionDate"`
      LastInstructionDate *string `xml:"LastInstructionDate,omitempty" json:"LastInstructionDate"`
      GraduationDate *GraduationDateType `xml:"GraduationDate,omitempty" json:"GraduationDate"`
      InstructionalMinutes *int `xml:"InstructionalMinutes" json:"InstructionalMinutes"`
      MinutesPerDay *int `xml:"MinutesPerDay" json:"MinutesPerDay"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    