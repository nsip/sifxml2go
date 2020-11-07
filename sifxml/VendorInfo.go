package sifxml

import _ "github.com/creasty/defaults"


type VendorInfos []VendorInfo

    type VendorInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      Name string `xml:"Name" json:"Name"`
      ContactInfo ContactInfoType `xml:"ContactInfo,omitempty" json:"ContactInfo"`
      CustomerId string `xml:"CustomerId,omitempty" json:"CustomerId"`
      ABN string `xml:"ABN,omitempty" json:"ABN"`
      RegisteredForGST AUCodeSetsYesOrNoCategoryType `xml:"RegisteredForGST,omitempty" json:"RegisteredForGST"`
      PaymentTerms string `xml:"PaymentTerms,omitempty" json:"PaymentTerms"`
      BPay string `xml:"BPay,omitempty" json:"BPay"`
      BSB string `xml:"BSB,omitempty" json:"BSB"`
      AccountNumber string `xml:"AccountNumber,omitempty" json:"AccountNumber"`
      AccountName string `xml:"AccountName,omitempty" json:"AccountName"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    