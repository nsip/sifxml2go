package sifxml

import _ "github.com/creasty/defaults"


type FinancialAccounts []FinancialAccount

    type FinancialAccount struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      ParentAccountRefId string `xml:"ParentAccountRefId,omitempty" json:"ParentAccountRefId"`
      ChargedLocationInfoRefId string `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId"`
      AccountNumber string `xml:"AccountNumber" json:"AccountNumber"`
      Name string `xml:"Name" json:"Name"`
      Description string `xml:"Description,omitempty" json:"Description"`
      ClassType string `xml:"ClassType" json:"ClassType"`
      AccountCode string `xml:"AccountCode,omitempty" json:"AccountCode"`
      CreationDate string `xml:"CreationDate" json:"CreationDate"`
      CreationTime string `xml:"CreationTime" json:"CreationTime"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    