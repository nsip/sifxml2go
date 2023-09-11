package sifxml



    type Invoices struct {
      invoices `json:"Invoices"`
    }

    type invoices struct {
      Invoice []invoice `json:"Invoice"`
    }

    type Invoice struct {
      invoice `xml:"Invoice" json:"Invoice"`
     }

     type invoice struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      InvoicedEntity *Invoice_InvoicedEntity `xml:"InvoicedEntity" json:"InvoicedEntity"`
      FormNumber *LocalIdType `xml:"FormNumber,omitempty" json:"FormNumber,omitempty"`
      BillingDate *String `xml:"BillingDate" json:"BillingDate"`
      TransactionDescription *String `xml:"TransactionDescription" json:"TransactionDescription"`
      BilledAmount *DebitOrCreditAmountType `xml:"BilledAmount" json:"BilledAmount"`
      Ledger *String `xml:"Ledger" json:"Ledger"`
      ChargedLocationInfoRefId *String `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId,omitempty"`
      NetAmount *MonetaryAmountType `xml:"NetAmount,omitempty" json:"NetAmount,omitempty"`
      TaxRate *Float `xml:"TaxRate,omitempty" json:"TaxRate,omitempty"`
      TaxType *String `xml:"TaxType,omitempty" json:"TaxType,omitempty"`
      TaxAmount *MonetaryAmountType `xml:"TaxAmount,omitempty" json:"TaxAmount,omitempty"`
      CreatedBy *String `xml:"CreatedBy,omitempty" json:"CreatedBy,omitempty"`
      ApprovedBy *String `xml:"ApprovedBy,omitempty" json:"ApprovedBy,omitempty"`
      ItemDetail *String `xml:"ItemDetail,omitempty" json:"ItemDetail,omitempty"`
      DueDate *String `xml:"DueDate,omitempty" json:"DueDate,omitempty"`
      FinancialAccountRefIdList *FinancialAccountRefIdListType `xml:"FinancialAccountRefIdList,omitempty" json:"FinancialAccountRefIdList,omitempty"`
      AccountCodeList *AccountCodeListType `xml:"AccountCodeList,omitempty" json:"AccountCodeList,omitempty"`
      AccountingPeriod *LocalIdType `xml:"AccountingPeriod,omitempty" json:"AccountingPeriod,omitempty"`
      RelatedPurchaseOrderRefId *String `xml:"RelatedPurchaseOrderRefId,omitempty" json:"RelatedPurchaseOrderRefId,omitempty"`
      PurchasingItems *PurchasingItemsType `xml:"PurchasingItems,omitempty" json:"PurchasingItems,omitempty"`
      Voluntary *AUCodeSetsYesOrNoCategoryType `xml:"Voluntary,omitempty" json:"Voluntary,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    
type Invoice_InvoicedEntity struct {
  invoice_invoicedentity `xml:"Invoice_InvoicedEntity" json:"Invoice_InvoicedEntity"`
}

type invoice_invoicedentity struct {

      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}

