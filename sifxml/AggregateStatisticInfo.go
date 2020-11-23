package sifxml


type AggregateStatisticInfos []AggregateStatisticInfo

    type AggregateStatisticInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StatisticName *String `xml:"StatisticName" json:"StatisticName"`
      CalculationRule *AggregateStatisticInfo_CalculationRule
      ApprovalDate *String `xml:"ApprovalDate,omitempty" json:"ApprovalDate"`
      ExpirationDate *String `xml:"ExpirationDate,omitempty" json:"ExpirationDate"`
      ExclusionRules *ExclusionRulesType `xml:"ExclusionRules,omitempty" json:"ExclusionRules"`
      Source *String `xml:"Source,omitempty" json:"Source"`
      EffectiveDate *String `xml:"EffectiveDate,omitempty" json:"EffectiveDate"`
      DiscontinueDate *String `xml:"DiscontinueDate,omitempty" json:"DiscontinueDate"`
      Location *LocationType `xml:"Location,omitempty" json:"Location"`
      Measure *String `xml:"Measure,omitempty" json:"Measure"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type AggregateStatisticInfo_CalculationRule struct {
      Type *String `xml:"Type,attr" json:"Type"`
      Value *String `xml:",chardata" json:"value"`
}
