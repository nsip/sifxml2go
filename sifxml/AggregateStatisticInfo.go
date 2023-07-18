package sifxml


type AggregateStatisticInfos []AggregateStatisticInfo

    type AggregateStatisticInfo struct {
  aggregatestatisticinfo `xml:"AggregateStatisticInfo" json:"AggregateStatisticInfo"`
}

type aggregatestatisticinfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StatisticName *String `xml:"StatisticName" json:"StatisticName"`
      CalculationRule *CalculationRuleType `xml:"CalculationRule,omitempty" json:"CalculationRule,omitempty"`
      ApprovalDate *String `xml:"ApprovalDate,omitempty" json:"ApprovalDate,omitempty"`
      ExpirationDate *String `xml:"ExpirationDate,omitempty" json:"ExpirationDate,omitempty"`
      ExclusionRules *ExclusionRulesType `xml:"ExclusionRules,omitempty" json:"ExclusionRules,omitempty"`
      Source *String `xml:"Source,omitempty" json:"Source,omitempty"`
      EffectiveDate *String `xml:"EffectiveDate,omitempty" json:"EffectiveDate,omitempty"`
      DiscontinueDate *String `xml:"DiscontinueDate,omitempty" json:"DiscontinueDate,omitempty"`
      Location *LocationType `xml:"Location,omitempty" json:"Location,omitempty"`
      Measure *String `xml:"Measure,omitempty" json:"Measure,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

