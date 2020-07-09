package sifxml


type MarkValueInfos []MarkValueInfo

    type MarkValueInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      Name *string `xml:"Name,omitempty" json:"Name"`
      PercentageMinimum *float64 `xml:"PercentageMinimum,omitempty" json:"PercentageMinimum"`
      PercentageMaximum *float64 `xml:"PercentageMaximum,omitempty" json:"PercentageMaximum"`
      PercentagePassingGrade *float64 `xml:"PercentagePassingGrade,omitempty" json:"PercentagePassingGrade"`
      NumericPrecision *int `xml:"NumericPrecision,omitempty" json:"NumericPrecision"`
      NumericScale *int `xml:"NumericScale,omitempty" json:"NumericScale"`
      NumericLow *float64 `xml:"NumericLow,omitempty" json:"NumericLow"`
      NumericHigh *float64 `xml:"NumericHigh,omitempty" json:"NumericHigh"`
      NumericPassingGrade *float64 `xml:"NumericPassingGrade,omitempty" json:"NumericPassingGrade"`
      ValidLetterMarkList *ValidLetterMarkListType `xml:"ValidLetterMarkList,omitempty" json:"ValidLetterMarkList"`
      Narrative *string `xml:"Narrative,omitempty" json:"Narrative"`
      NarrativeMaximumSize *int `xml:"NarrativeMaximumSize,omitempty" json:"NarrativeMaximumSize"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    