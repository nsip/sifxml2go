package sifxml


    type Debtor struct {
        RefId RefIdType `xml:"RefId,attr"`
      LocalId LocalIdType `xml:"LocalId"`
      BilledEntity Debtor_BilledEntity `xml:"BilledEntity"`
      AddressList AddressListType `xml:"AddressList"`
      BillingName string `xml:"BillingName"`
      BillingNote string `xml:"BillingNote"`
      Discount string `xml:"Discount"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    type Debtor_BilledEntity struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr"`
      Value string `xml:",chardata"`
}
