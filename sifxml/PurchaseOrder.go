package sifxml



    type PurchaseOrders struct {
      purchaseorders `json:"PurchaseOrders"`
    }

    type purchaseorders struct {
      PurchaseOrder []purchaseorder `json:"PurchaseOrder"`
    }

    type PurchaseOrder struct {
      purchaseorder `xml:"PurchaseOrder" json:"PurchaseOrder"`
     }

     type purchaseorder struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      FormNumber *String `xml:"FormNumber" json:"FormNumber"`
      VendorInfoRefId *String `xml:"VendorInfoRefId" json:"VendorInfoRefId"`
      ChargedLocationInfoRefId *String `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId,omitempty"`
      EmployeePersonalRefId *String `xml:"EmployeePersonalRefId,omitempty" json:"EmployeePersonalRefId,omitempty"`
      PurchasingItems *PurchasingItemsType `xml:"PurchasingItems" json:"PurchasingItems"`
      CreationDate *String `xml:"CreationDate,omitempty" json:"CreationDate,omitempty"`
      TaxRate *Float `xml:"TaxRate,omitempty" json:"TaxRate,omitempty"`
      TaxAmount *MonetaryAmountType `xml:"TaxAmount,omitempty" json:"TaxAmount,omitempty"`
      TotalAmount *MonetaryAmountType `xml:"TotalAmount,omitempty" json:"TotalAmount,omitempty"`
      UpdateDate *String `xml:"UpdateDate,omitempty" json:"UpdateDate,omitempty"`
      FullyDelivered *AUCodeSetsYesOrNoCategoryType `xml:"FullyDelivered,omitempty" json:"FullyDelivered,omitempty"`
      OriginalPurchaseOrderRefId *String `xml:"OriginalPurchaseOrderRefId,omitempty" json:"OriginalPurchaseOrderRefId,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

