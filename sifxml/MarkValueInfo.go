package sifxml


type MarkValueInfos []MarkValueInfo

    type MarkValueInfo struct {
  markvalueinfo `xml:"MarkValueInfo" json:"MarkValueInfo"`
}

type markvalueinfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels,omitempty"`
      Name *String `xml:"Name" json:"Name"`
      PercentageMinimum *Float `xml:"PercentageMinimum,omitempty" json:"PercentageMinimum,omitempty"`
      PercentageMaximum *Float `xml:"PercentageMaximum,omitempty" json:"PercentageMaximum,omitempty"`
      PercentagePassingGrade *Float `xml:"PercentagePassingGrade,omitempty" json:"PercentagePassingGrade,omitempty"`
      NumericPrecision *Int `xml:"NumericPrecision,omitempty" json:"NumericPrecision,omitempty"`
      NumericScale *Int `xml:"NumericScale,omitempty" json:"NumericScale,omitempty"`
      NumericLow *Float `xml:"NumericLow,omitempty" json:"NumericLow,omitempty"`
      NumericHigh *Float `xml:"NumericHigh,omitempty" json:"NumericHigh,omitempty"`
      NumericPassingGrade *Float `xml:"NumericPassingGrade,omitempty" json:"NumericPassingGrade,omitempty"`
      ValidLetterMarkList *ValidLetterMarkListType `xml:"ValidLetterMarkList,omitempty" json:"ValidLetterMarkList,omitempty"`
      Narrative *String `xml:"Narrative" json:"Narrative"`
      NarrativeMaximumSize *Int `xml:"NarrativeMaximumSize,omitempty" json:"NarrativeMaximumSize,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    