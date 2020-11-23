package sifxml


type Debtors []Debtor

    type Debtor struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      BilledEntity *Debtor_BilledEntity `xml:"BilledEntity" json:"BilledEntity"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      BillingName *String `xml:"BillingName,omitempty" json:"BillingName"`
      BillingNote *String `xml:"BillingNote,omitempty" json:"BillingNote"`
      Discount *String `xml:"Discount,omitempty" json:"Discount"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type Debtor_BilledEntity struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
