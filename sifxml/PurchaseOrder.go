package sifxml


type PurchaseOrders []PurchaseOrder

    type PurchaseOrder struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      FormNumber *String `xml:"FormNumber" json:"FormNumber"`
      VendorInfoRefId *String `xml:"VendorInfoRefId" json:"VendorInfoRefId"`
      ChargedLocationInfoRefId *String `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId"`
      EmployeePersonalRefId *String `xml:"EmployeePersonalRefId,omitempty" json:"EmployeePersonalRefId"`
      PurchasingItems *PurchasingItemsType `xml:"PurchasingItems" json:"PurchasingItems"`
      CreationDate *String `xml:"CreationDate,omitempty" json:"CreationDate"`
      TaxRate *Float `xml:"TaxRate,omitempty" json:"TaxRate"`
      TaxAmount *MonetaryAmountType `xml:"TaxAmount,omitempty" json:"TaxAmount"`
      TotalAmount *MonetaryAmountType `xml:"TotalAmount,omitempty" json:"TotalAmount"`
      UpdateDate *String `xml:"UpdateDate,omitempty" json:"UpdateDate"`
      FullyDelivered *AUCodeSetsYesOrNoCategoryType `xml:"FullyDelivered,omitempty" json:"FullyDelivered"`
      OriginalPurchaseOrderRefId *String `xml:"OriginalPurchaseOrderRefId,omitempty" json:"OriginalPurchaseOrderRefId"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    