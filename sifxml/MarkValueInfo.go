package sifxml

import _ "github.com/creasty/defaults"


type MarkValueInfos []MarkValueInfo

    type MarkValueInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId" json:"SchoolInfoRefId"`
      YearLevels YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      Name string `xml:"Name" json:"Name"`
      PercentageMinimum float64 `default:"-2147483648" xml:"PercentageMinimum" json:"PercentageMinimum"`
      PercentageMaximum float64 `default:"-2147483648" xml:"PercentageMaximum" json:"PercentageMaximum"`
      PercentagePassingGrade float64 `default:"-2147483648" xml:"PercentagePassingGrade" json:"PercentagePassingGrade"`
      NumericPrecision int `default:"-2147483648" xml:"NumericPrecision" json:"NumericPrecision"`
      NumericScale int `default:"-2147483648" xml:"NumericScale" json:"NumericScale"`
      NumericLow float64 `default:"-2147483648" xml:"NumericLow" json:"NumericLow"`
      NumericHigh float64 `default:"-2147483648" xml:"NumericHigh" json:"NumericHigh"`
      NumericPassingGrade float64 `default:"-2147483648" xml:"NumericPassingGrade" json:"NumericPassingGrade"`
      ValidLetterMarkList ValidLetterMarkListType `xml:"ValidLetterMarkList" json:"ValidLetterMarkList"`
      Narrative string `xml:"Narrative" json:"Narrative"`
      NarrativeMaximumSize int `default:"-2147483648" xml:"NarrativeMaximumSize" json:"NarrativeMaximumSize"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    