package sifxml


type FinancialAccounts []FinancialAccount

    type FinancialAccount struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      ParentAccountRefId *String `xml:"ParentAccountRefId,omitempty" json:"ParentAccountRefId"`
      ChargedLocationInfoRefId *String `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId"`
      AccountNumber *String `xml:"AccountNumber" json:"AccountNumber"`
      Name *String `xml:"Name" json:"Name"`
      Description *String `xml:"Description,omitempty" json:"Description"`
      ClassType *String `xml:"ClassType" json:"ClassType"`
      AccountCode *String `xml:"AccountCode,omitempty" json:"AccountCode"`
      CreationDate *String `xml:"CreationDate" json:"CreationDate"`
      CreationTime *String `xml:"CreationTime" json:"CreationTime"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    