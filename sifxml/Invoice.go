package sifxml


    type Invoice struct {
        RefId RefIdType `xml:"RefId,attr"`
      LocalId LocalIdType `xml:"LocalId"`
      InvoicedEntity Invoice_InvoicedEntity `xml:"InvoicedEntity"`
      FormNumber LocalIdType `xml:"FormNumber"`
      BillingDate string `xml:"BillingDate"`
      TransactionDescription string `xml:"TransactionDescription"`
      BilledAmount DebitOrCreditAmountType `xml:"BilledAmount"`
      Ledger string `xml:"Ledger"`
      ChargedLocationInfoRefId string `xml:"ChargedLocationInfoRefId"`
      NetAmount MonetaryAmountType `xml:"NetAmount"`
      TaxRate string `xml:"TaxRate"`
      TaxType string `xml:"TaxType"`
      TaxAmount MonetaryAmountType `xml:"TaxAmount"`
      CreatedBy string `xml:"CreatedBy"`
      ApprovedBy string `xml:"ApprovedBy"`
      ItemDetail string `xml:"ItemDetail"`
      DueDate string `xml:"DueDate"`
      FinancialAccountRefIdList FinancialAccountRefIdListType `xml:"FinancialAccountRefIdList"`
      AccountCodeList AccountCodeListType `xml:"AccountCodeList"`
      AccountingPeriod LocalIdType `xml:"AccountingPeriod"`
      RelatedPurchaseOrderRefId string `xml:"RelatedPurchaseOrderRefId"`
      PurchasingItems PurchasingItemsType `xml:"PurchasingItems"`
      Voluntary string `xml:"Voluntary"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    type Invoice_InvoicedEntity struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr"`
      Value string `xml:",chardata"`
}
