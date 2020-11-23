package sifxml


type TermInfos []TermInfo

    type TermInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      StartDate *String `xml:"StartDate" json:"StartDate"`
      EndDate *String `xml:"EndDate" json:"EndDate"`
      Description *String `xml:"Description,omitempty" json:"Description"`
      RelativeDuration *Float `xml:"RelativeDuration,omitempty" json:"RelativeDuration"`
      TermCode *String `xml:"TermCode,omitempty" json:"TermCode"`
      Track *String `xml:"Track,omitempty" json:"Track"`
      TermSpan *AUCodeSetsSessionTypeType `xml:"TermSpan,omitempty" json:"TermSpan"`
      MarkingTerm *AUCodeSetsYesOrNoCategoryType `xml:"MarkingTerm,omitempty" json:"MarkingTerm"`
      SchedulingTerm *AUCodeSetsYesOrNoCategoryType `xml:"SchedulingTerm,omitempty" json:"SchedulingTerm"`
      AttendanceTerm *AUCodeSetsYesOrNoCategoryType `xml:"AttendanceTerm,omitempty" json:"AttendanceTerm"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    