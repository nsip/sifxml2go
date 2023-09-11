package sifxml



    type Debtors struct {
      debtors `json:"Debtors"`
    }

    type debtors struct {
      Debtor []debtor `json:"Debtor"`
    }

    type Debtor struct {
      debtor `xml:"Debtor" json:"Debtor"`
     }

     type debtor struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      BilledEntity *Debtor_BilledEntity `xml:"BilledEntity" json:"BilledEntity"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList,omitempty"`
      BillingName *String `xml:"BillingName,omitempty" json:"BillingName,omitempty"`
      BillingNote *String `xml:"BillingNote,omitempty" json:"BillingNote,omitempty"`
      Discount *String `xml:"Discount,omitempty" json:"Discount,omitempty"`
      BSB *String `xml:"BSB,omitempty" json:"BSB,omitempty"`
      AccountNumber *String `xml:"AccountNumber,omitempty" json:"AccountNumber,omitempty"`
      AccountName *String `xml:"AccountName,omitempty" json:"AccountName,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    
type Debtor_BilledEntity struct {
  debtor_billedentity `xml:"Debtor_BilledEntity" json:"Debtor_BilledEntity"`
}

type debtor_billedentity struct {

      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}

