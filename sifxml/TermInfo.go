package sifxml

import _ "github.com/creasty/defaults"


type TermInfos []TermInfo

    type TermInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      StartDate string `xml:"StartDate" json:"StartDate"`
      EndDate string `xml:"EndDate" json:"EndDate"`
      Description string `xml:"Description,omitempty" json:"Description"`
      RelativeDuration float64 `default:"-2147483648" xml:"RelativeDuration" json:"RelativeDuration"`
      TermCode string `xml:"TermCode,omitempty" json:"TermCode"`
      Track string `xml:"Track,omitempty" json:"Track"`
      TermSpan AUCodeSetsSessionTypeType `xml:"TermSpan,omitempty" json:"TermSpan"`
      MarkingTerm AUCodeSetsYesOrNoCategoryType `xml:"MarkingTerm,omitempty" json:"MarkingTerm"`
      SchedulingTerm AUCodeSetsYesOrNoCategoryType `xml:"SchedulingTerm,omitempty" json:"SchedulingTerm"`
      AttendanceTerm AUCodeSetsYesOrNoCategoryType `xml:"AttendanceTerm,omitempty" json:"AttendanceTerm"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    