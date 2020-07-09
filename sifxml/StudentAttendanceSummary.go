package sifxml


type StudentAttendanceSummarys []StudentAttendanceSummary

    type StudentAttendanceSummary struct {
        StudentAttendanceSummaryRefId *string `xml:"StudentAttendanceSummaryRefId,attr" json:"StudentAttendanceSummaryRefId"`
      StudentPersonalRefId *string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      StartDate *string `xml:"StartDate,omitempty" json:"StartDate"`
      EndDate *string `xml:"EndDate,omitempty" json:"EndDate"`
      StartDay *int `xml:"StartDay,omitempty" json:"StartDay"`
      EndDay *int `xml:"EndDay,omitempty" json:"EndDay"`
      FTE *float64 `xml:"FTE,omitempty" json:"FTE"`
      DaysAttended *float64 `xml:"DaysAttended,omitempty" json:"DaysAttended"`
      ExcusedAbsences *float64 `xml:"ExcusedAbsences,omitempty" json:"ExcusedAbsences"`
      UnexcusedAbsences *float64 `xml:"UnexcusedAbsences,omitempty" json:"UnexcusedAbsences"`
      DaysTardy *float64 `xml:"DaysTardy,omitempty" json:"DaysTardy"`
      DaysInMembership *float64 `xml:"DaysInMembership,omitempty" json:"DaysInMembership"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    