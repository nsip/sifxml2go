package sifxml


type PaymentReceipts []PaymentReceipt

    type PaymentReceipt struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      TransactionType *string `xml:"TransactionType" json:"TransactionType"`
      InvoiceRefId *string `xml:"InvoiceRefId,omitempty" json:"InvoiceRefId"`
      PaymentReceiptLineList *PaymentReceiptLineListType `xml:"PaymentReceiptLineList,omitempty" json:"PaymentReceiptLineList"`
      VendorInfoRefId *string `xml:"VendorInfoRefId,omitempty" json:"VendorInfoRefId"`
      DebtorRefId *string `xml:"DebtorRefId,omitempty" json:"DebtorRefId"`
      ChargedLocationInfoRefId *string `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId"`
      TransactionDate *string `xml:"TransactionDate" json:"TransactionDate"`
      TransactionAmount *DebitOrCreditAmountType `xml:"TransactionAmount" json:"TransactionAmount"`
      ReceivedTransactionId *string `xml:"ReceivedTransactionId" json:"ReceivedTransactionId"`
      FinancialAccountRefIdList *FinancialAccountRefIdListType `xml:"FinancialAccountRefIdList,omitempty" json:"FinancialAccountRefIdList"`
      AccountCodeList *AccountCodeListType `xml:"AccountCodeList,omitempty" json:"AccountCodeList"`
      TransactionDescription *string `xml:"TransactionDescription,omitempty" json:"TransactionDescription"`
      TaxRate *float64 `xml:"TaxRate" json:"TaxRate"`
      TaxAmount *MonetaryAmountType `xml:"TaxAmount,omitempty" json:"TaxAmount"`
      TransactionMethod *string `xml:"TransactionMethod,omitempty" json:"TransactionMethod"`
      ChequeNumber *string `xml:"ChequeNumber,omitempty" json:"ChequeNumber"`
      TransactionNote *string `xml:"TransactionNote,omitempty" json:"TransactionNote"`
      AccountingPeriod *LocalIdType `xml:"AccountingPeriod,omitempty" json:"AccountingPeriod"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    