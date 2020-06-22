package sifxml


    type PurchaseOrder struct {
        RefId RefIdType `xml:"RefId,attr"`
      LocalId LocalIdType `xml:"LocalId"`
      FormNumber string `xml:"FormNumber"`
      VendorInfoRefId string `xml:"VendorInfoRefId"`
      ChargedLocationInfoRefId string `xml:"ChargedLocationInfoRefId"`
      EmployeePersonalRefId string `xml:"EmployeePersonalRefId"`
      PurchasingItems PurchasingItemsType `xml:"PurchasingItems"`
      CreationDate string `xml:"CreationDate"`
      TaxRate string `xml:"TaxRate"`
      TaxAmount MonetaryAmountType `xml:"TaxAmount"`
      TotalAmount MonetaryAmountType `xml:"TotalAmount"`
      UpdateDate string `xml:"UpdateDate"`
      FullyDelivered string `xml:"FullyDelivered"`
      OriginalPurchaseOrderRefId string `xml:"OriginalPurchaseOrderRefId"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    