package sifxml


type VendorInfos []VendorInfo

    type VendorInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      Name *String `xml:"Name" json:"Name"`
      ContactInfo *ContactInfoType `xml:"ContactInfo,omitempty" json:"ContactInfo"`
      CustomerId *String `xml:"CustomerId,omitempty" json:"CustomerId"`
      ABN *String `xml:"ABN,omitempty" json:"ABN"`
      RegisteredForGST *AUCodeSetsYesOrNoCategoryType `xml:"RegisteredForGST,omitempty" json:"RegisteredForGST"`
      PaymentTerms *String `xml:"PaymentTerms,omitempty" json:"PaymentTerms"`
      BPay *String `xml:"BPay,omitempty" json:"BPay"`
      BSB *String `xml:"BSB,omitempty" json:"BSB"`
      AccountNumber *String `xml:"AccountNumber,omitempty" json:"AccountNumber"`
      AccountName *String `xml:"AccountName,omitempty" json:"AccountName"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    