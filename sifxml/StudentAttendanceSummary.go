package sifxml

import _ "github.com/creasty/defaults"


type StudentAttendanceSummarys []StudentAttendanceSummary

    type StudentAttendanceSummary struct {
        StudentAttendanceSummaryRefId string `xml:"StudentAttendanceSummaryRefId,attr" json:"StudentAttendanceSummaryRefId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId" json:"StudentPersonalRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      StartDate string `xml:"StartDate" json:"StartDate"`
      EndDate string `xml:"EndDate" json:"EndDate"`
      StartDay int `default:"-2147483648" xml:"StartDay" json:"StartDay"`
      EndDay int `default:"-2147483648" xml:"EndDay" json:"EndDay"`
      FTE float64 `default:"-2147483648" xml:"FTE" json:"FTE"`
      DaysAttended float64 `xml:"DaysAttended" json:"DaysAttended"`
      ExcusedAbsences float64 `xml:"ExcusedAbsences" json:"ExcusedAbsences"`
      UnexcusedAbsences float64 `xml:"UnexcusedAbsences" json:"UnexcusedAbsences"`
      DaysTardy float64 `default:"-2147483648" xml:"DaysTardy" json:"DaysTardy"`
      DaysInMembership float64 `xml:"DaysInMembership" json:"DaysInMembership"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    