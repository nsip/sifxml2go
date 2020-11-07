package sifxml


type Invoices []Invoice

    type Invoice struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      InvoicedEntity *Invoice_InvoicedEntity `xml:"InvoicedEntity" json:"InvoicedEntity"`
      FormNumber *LocalIdType `xml:"FormNumber,omitempty" json:"FormNumber"`
      BillingDate *string `xml:"BillingDate" json:"BillingDate"`
      TransactionDescription *string `xml:"TransactionDescription" json:"TransactionDescription"`
      BilledAmount *DebitOrCreditAmountType `xml:"BilledAmount" json:"BilledAmount"`
      Ledger *string `xml:"Ledger" json:"Ledger"`
      ChargedLocationInfoRefId *string `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId"`
      NetAmount *MonetaryAmountType `xml:"NetAmount,omitempty" json:"NetAmount"`
      TaxRate *float64 `xml:"TaxRate" json:"TaxRate"`
      TaxType *string `xml:"TaxType,omitempty" json:"TaxType"`
      TaxAmount *MonetaryAmountType `xml:"TaxAmount,omitempty" json:"TaxAmount"`
      CreatedBy *string `xml:"CreatedBy,omitempty" json:"CreatedBy"`
      ApprovedBy *string `xml:"ApprovedBy,omitempty" json:"ApprovedBy"`
      ItemDetail *string `xml:"ItemDetail,omitempty" json:"ItemDetail"`
      DueDate *string `xml:"DueDate,omitempty" json:"DueDate"`
      FinancialAccountRefIdList *FinancialAccountRefIdListType `xml:"FinancialAccountRefIdList,omitempty" json:"FinancialAccountRefIdList"`
      AccountCodeList *AccountCodeListType `xml:"AccountCodeList,omitempty" json:"AccountCodeList"`
      AccountingPeriod *LocalIdType `xml:"AccountingPeriod,omitempty" json:"AccountingPeriod"`
      RelatedPurchaseOrderRefId *string `xml:"RelatedPurchaseOrderRefId,omitempty" json:"RelatedPurchaseOrderRefId"`
      PurchasingItems *PurchasingItemsType `xml:"PurchasingItems,omitempty" json:"PurchasingItems"`
      Voluntary *AUCodeSetsYesOrNoCategoryType `xml:"Voluntary,omitempty" json:"Voluntary"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type Invoice_InvoicedEntity struct {
      SIF_RefObject *string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *string `xml:",chardata" json:"value"`
}
