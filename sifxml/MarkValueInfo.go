package sifxml


    type MarkValueInfo struct {
        RefId RefIdType `xml:"RefId,attr"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      YearLevels YearLevelsType `xml:"YearLevels"`
      Name string `xml:"Name"`
      PercentageMinimum string `xml:"PercentageMinimum"`
      PercentageMaximum string `xml:"PercentageMaximum"`
      PercentagePassingGrade string `xml:"PercentagePassingGrade"`
      NumericPrecision string `xml:"NumericPrecision"`
      NumericScale string `xml:"NumericScale"`
      NumericLow string `xml:"NumericLow"`
      NumericHigh string `xml:"NumericHigh"`
      NumericPassingGrade string `xml:"NumericPassingGrade"`
      ValidLetterMarkList ValidLetterMarkListType `xml:"ValidLetterMarkList"`
      Narrative string `xml:"Narrative"`
      NarrativeMaximumSize string `xml:"NarrativeMaximumSize"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    