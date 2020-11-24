package sifxml


type TermInfos []TermInfo

    type TermInfo struct {
  terminfo `xml:"TermInfo" json:"TermInfo"`
}

type terminfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      StartDate *String `xml:"StartDate" json:"StartDate"`
      EndDate *String `xml:"EndDate" json:"EndDate"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      RelativeDuration *Float `xml:"RelativeDuration,omitempty" json:"RelativeDuration,omitempty"`
      TermCode *String `xml:"TermCode,omitempty" json:"TermCode,omitempty"`
      Track *String `xml:"Track,omitempty" json:"Track,omitempty"`
      TermSpan *AUCodeSetsSessionTypeType `xml:"TermSpan,omitempty" json:"TermSpan,omitempty"`
      MarkingTerm *AUCodeSetsYesOrNoCategoryType `xml:"MarkingTerm,omitempty" json:"MarkingTerm,omitempty"`
      SchedulingTerm *AUCodeSetsYesOrNoCategoryType `xml:"SchedulingTerm,omitempty" json:"SchedulingTerm,omitempty"`
      AttendanceTerm *AUCodeSetsYesOrNoCategoryType `xml:"AttendanceTerm,omitempty" json:"AttendanceTerm,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    