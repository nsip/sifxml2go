package sifxml


type MarkValueInfos []MarkValueInfo

    type MarkValueInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      Name *String `xml:"Name" json:"Name"`
      PercentageMinimum *Float `xml:"PercentageMinimum,omitempty" json:"PercentageMinimum"`
      PercentageMaximum *Float `xml:"PercentageMaximum,omitempty" json:"PercentageMaximum"`
      PercentagePassingGrade *Float `xml:"PercentagePassingGrade,omitempty" json:"PercentagePassingGrade"`
      NumericPrecision *Int `xml:"NumericPrecision,omitempty" json:"NumericPrecision"`
      NumericScale *Int `xml:"NumericScale,omitempty" json:"NumericScale"`
      NumericLow *Float `xml:"NumericLow,omitempty" json:"NumericLow"`
      NumericHigh *Float `xml:"NumericHigh,omitempty" json:"NumericHigh"`
      NumericPassingGrade *Float `xml:"NumericPassingGrade,omitempty" json:"NumericPassingGrade"`
      ValidLetterMarkList *ValidLetterMarkListType `xml:"ValidLetterMarkList,omitempty" json:"ValidLetterMarkList"`
      Narrative *String `xml:"Narrative" json:"Narrative"`
      NarrativeMaximumSize *Int `xml:"NarrativeMaximumSize,omitempty" json:"NarrativeMaximumSize"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    