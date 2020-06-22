package sifxml


    type FinancialAccount struct {
        RefId RefIdType `xml:"RefId,attr"`
      LocalId LocalIdType `xml:"LocalId"`
      ParentAccountRefId string `xml:"ParentAccountRefId"`
      ChargedLocationInfoRefId string `xml:"ChargedLocationInfoRefId"`
      AccountNumber string `xml:"AccountNumber"`
      Name string `xml:"Name"`
      Description string `xml:"Description"`
      ClassType string `xml:"ClassType"`
      AccountCode string `xml:"AccountCode"`
      CreationDate string `xml:"CreationDate"`
      CreationTime string `xml:"CreationTime"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    