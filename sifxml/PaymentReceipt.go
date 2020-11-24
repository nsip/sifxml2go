package sifxml


type PaymentReceipts []PaymentReceipt

    type PaymentReceipt struct {
  paymentreceipt `xml:"PaymentReceipt" json:"PaymentReceipt"`
}

type paymentreceipt struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      TransactionType *String `xml:"TransactionType" json:"TransactionType"`
      InvoiceRefId *String `xml:"InvoiceRefId,omitempty" json:"InvoiceRefId,omitempty"`
      PaymentReceiptLineList *PaymentReceiptLineListType `xml:"PaymentReceiptLineList,omitempty" json:"PaymentReceiptLineList,omitempty"`
      VendorInfoRefId *String `xml:"VendorInfoRefId,omitempty" json:"VendorInfoRefId,omitempty"`
      DebtorRefId *String `xml:"DebtorRefId,omitempty" json:"DebtorRefId,omitempty"`
      ChargedLocationInfoRefId *String `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId,omitempty"`
      TransactionDate *String `xml:"TransactionDate" json:"TransactionDate"`
      TransactionAmount *DebitOrCreditAmountType `xml:"TransactionAmount" json:"TransactionAmount"`
      ReceivedTransactionId *String `xml:"ReceivedTransactionId" json:"ReceivedTransactionId"`
      FinancialAccountRefIdList *FinancialAccountRefIdListType `xml:"FinancialAccountRefIdList,omitempty" json:"FinancialAccountRefIdList,omitempty"`
      AccountCodeList *AccountCodeListType `xml:"AccountCodeList,omitempty" json:"AccountCodeList,omitempty"`
      TransactionDescription *String `xml:"TransactionDescription,omitempty" json:"TransactionDescription,omitempty"`
      TaxRate *Float `xml:"TaxRate,omitempty" json:"TaxRate,omitempty"`
      TaxAmount *MonetaryAmountType `xml:"TaxAmount,omitempty" json:"TaxAmount,omitempty"`
      TransactionMethod *String `xml:"TransactionMethod,omitempty" json:"TransactionMethod,omitempty"`
      ChequeNumber *String `xml:"ChequeNumber,omitempty" json:"ChequeNumber,omitempty"`
      TransactionNote *String `xml:"TransactionNote,omitempty" json:"TransactionNote,omitempty"`
      AccountingPeriod *LocalIdType `xml:"AccountingPeriod,omitempty" json:"AccountingPeriod,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    