package sifxml


type VendorInfos []VendorInfo

    type VendorInfo struct {
  vendorinfo `xml:"VendorInfo" json:"VendorInfo"`
}

type vendorinfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      Name *String `xml:"Name" json:"Name"`
      ContactInfo *ContactInfoType `xml:"ContactInfo,omitempty" json:"ContactInfo,omitempty"`
      CustomerId *String `xml:"CustomerId,omitempty" json:"CustomerId,omitempty"`
      ABN *String `xml:"ABN,omitempty" json:"ABN,omitempty"`
      RegisteredForGST *AUCodeSetsYesOrNoCategoryType `xml:"RegisteredForGST,omitempty" json:"RegisteredForGST,omitempty"`
      PaymentTerms *String `xml:"PaymentTerms,omitempty" json:"PaymentTerms,omitempty"`
      BPay *String `xml:"BPay,omitempty" json:"BPay,omitempty"`
      BSB *String `xml:"BSB,omitempty" json:"BSB,omitempty"`
      AccountNumber *String `xml:"AccountNumber,omitempty" json:"AccountNumber,omitempty"`
      AccountName *String `xml:"AccountName,omitempty" json:"AccountName,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

