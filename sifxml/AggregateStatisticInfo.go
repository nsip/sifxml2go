package sifxml

import _ "github.com/creasty/defaults"


type AggregateStatisticInfos []AggregateStatisticInfo

    type AggregateStatisticInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      StatisticName string `xml:"StatisticName" json:"StatisticName"`
      CalculationRule AggregateStatisticInfo_CalculationRule
      ApprovalDate string `xml:"ApprovalDate,omitempty" json:"ApprovalDate"`
      ExpirationDate string `xml:"ExpirationDate,omitempty" json:"ExpirationDate"`
      ExclusionRules ExclusionRulesType `xml:"ExclusionRules,omitempty" json:"ExclusionRules"`
      Source string `xml:"Source,omitempty" json:"Source"`
      EffectiveDate string `xml:"EffectiveDate,omitempty" json:"EffectiveDate"`
      DiscontinueDate string `xml:"DiscontinueDate,omitempty" json:"DiscontinueDate"`
      Location LocationType `xml:"Location,omitempty" json:"Location"`
      Measure string `xml:"Measure,omitempty" json:"Measure"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type AggregateStatisticInfo_CalculationRule struct {
      Type string `xml:"Type,attr" json:"Type"`
      Value string `xml:",chardata" json:"value"`
}
