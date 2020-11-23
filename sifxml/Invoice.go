package sifxml


type Invoices []Invoice

    type Invoice struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      InvoicedEntity *Invoice_InvoicedEntity `xml:"InvoicedEntity" json:"InvoicedEntity"`
      FormNumber *LocalIdType `xml:"FormNumber,omitempty" json:"FormNumber"`
      BillingDate *String `xml:"BillingDate" json:"BillingDate"`
      TransactionDescription *String `xml:"TransactionDescription" json:"TransactionDescription"`
      BilledAmount *DebitOrCreditAmountType `xml:"BilledAmount" json:"BilledAmount"`
      Ledger *String `xml:"Ledger" json:"Ledger"`
      ChargedLocationInfoRefId *String `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId"`
      NetAmount *MonetaryAmountType `xml:"NetAmount,omitempty" json:"NetAmount"`
      TaxRate *Float `xml:"TaxRate,omitempty" json:"TaxRate"`
      TaxType *String `xml:"TaxType,omitempty" json:"TaxType"`
      TaxAmount *MonetaryAmountType `xml:"TaxAmount,omitempty" json:"TaxAmount"`
      CreatedBy *String `xml:"CreatedBy,omitempty" json:"CreatedBy"`
      ApprovedBy *String `xml:"ApprovedBy,omitempty" json:"ApprovedBy"`
      ItemDetail *String `xml:"ItemDetail,omitempty" json:"ItemDetail"`
      DueDate *String `xml:"DueDate,omitempty" json:"DueDate"`
      FinancialAccountRefIdList *FinancialAccountRefIdListType `xml:"FinancialAccountRefIdList,omitempty" json:"FinancialAccountRefIdList"`
      AccountCodeList *AccountCodeListType `xml:"AccountCodeList,omitempty" json:"AccountCodeList"`
      AccountingPeriod *LocalIdType `xml:"AccountingPeriod,omitempty" json:"AccountingPeriod"`
      RelatedPurchaseOrderRefId *String `xml:"RelatedPurchaseOrderRefId,omitempty" json:"RelatedPurchaseOrderRefId"`
      PurchasingItems *PurchasingItemsType `xml:"PurchasingItems,omitempty" json:"PurchasingItems"`
      Voluntary *AUCodeSetsYesOrNoCategoryType `xml:"Voluntary,omitempty" json:"Voluntary"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type Invoice_InvoicedEntity struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
