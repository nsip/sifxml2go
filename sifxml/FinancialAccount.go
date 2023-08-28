package sifxml


type FinancialAccounts []FinancialAccount

    type FinancialAccount struct {
  financialaccount `xml:"FinancialAccount"`
}

type financialaccount struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      ParentAccountRefId *String `xml:"ParentAccountRefId,omitempty" json:"ParentAccountRefId,omitempty"`
      ChargedLocationInfoRefId *String `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId,omitempty"`
      AccountNumber *String `xml:"AccountNumber" json:"AccountNumber"`
      Name *String `xml:"Name" json:"Name"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      ClassType *String `xml:"ClassType" json:"ClassType"`
      AccountCode *String `xml:"AccountCode,omitempty" json:"AccountCode,omitempty"`
      CreationDate *String `xml:"CreationDate" json:"CreationDate"`
      CreationTime *String `xml:"CreationTime" json:"CreationTime"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

