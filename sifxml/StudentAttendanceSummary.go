package sifxml


type StudentAttendanceSummarys []StudentAttendanceSummary

    type StudentAttendanceSummary struct {
  studentattendancesummary `xml:"StudentAttendanceSummary" json:"StudentAttendanceSummary"`
}

type studentattendancesummary struct {
        StudentAttendanceSummaryRefId *String `xml:"StudentAttendanceSummaryRefId,attr" json:"StudentAttendanceSummaryRefId"`
      StudentPersonalRefId *String `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      StartDate *String `xml:"StartDate" json:"StartDate"`
      EndDate *String `xml:"EndDate" json:"EndDate"`
      StartDay *Int `xml:"StartDay,omitempty" json:"StartDay,omitempty"`
      EndDay *Int `xml:"EndDay,omitempty" json:"EndDay,omitempty"`
      FTE *Float `xml:"FTE,omitempty" json:"FTE,omitempty"`
      DaysAttended *Float `xml:"DaysAttended" json:"DaysAttended"`
      ExcusedAbsences *Float `xml:"ExcusedAbsences" json:"ExcusedAbsences"`
      UnexcusedAbsences *Float `xml:"UnexcusedAbsences" json:"UnexcusedAbsences"`
      DaysTardy *Float `xml:"DaysTardy,omitempty" json:"DaysTardy,omitempty"`
      DaysInMembership *Float `xml:"DaysInMembership" json:"DaysInMembership"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    