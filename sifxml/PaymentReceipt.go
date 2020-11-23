package sifxml


type PaymentReceipts []PaymentReceipt

    type PaymentReceipt struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      TransactionType *String `xml:"TransactionType" json:"TransactionType"`
      InvoiceRefId *String `xml:"InvoiceRefId,omitempty" json:"InvoiceRefId"`
      PaymentReceiptLineList *PaymentReceiptLineListType `xml:"PaymentReceiptLineList,omitempty" json:"PaymentReceiptLineList"`
      VendorInfoRefId *String `xml:"VendorInfoRefId,omitempty" json:"VendorInfoRefId"`
      DebtorRefId *String `xml:"DebtorRefId,omitempty" json:"DebtorRefId"`
      ChargedLocationInfoRefId *String `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId"`
      TransactionDate *String `xml:"TransactionDate" json:"TransactionDate"`
      TransactionAmount *DebitOrCreditAmountType `xml:"TransactionAmount" json:"TransactionAmount"`
      ReceivedTransactionId *String `xml:"ReceivedTransactionId" json:"ReceivedTransactionId"`
      FinancialAccountRefIdList *FinancialAccountRefIdListType `xml:"FinancialAccountRefIdList,omitempty" json:"FinancialAccountRefIdList"`
      AccountCodeList *AccountCodeListType `xml:"AccountCodeList,omitempty" json:"AccountCodeList"`
      TransactionDescription *String `xml:"TransactionDescription,omitempty" json:"TransactionDescription"`
      TaxRate *Float `xml:"TaxRate,omitempty" json:"TaxRate"`
      TaxAmount *MonetaryAmountType `xml:"TaxAmount,omitempty" json:"TaxAmount"`
      TransactionMethod *String `xml:"TransactionMethod,omitempty" json:"TransactionMethod"`
      ChequeNumber *String `xml:"ChequeNumber,omitempty" json:"ChequeNumber"`
      TransactionNote *String `xml:"TransactionNote,omitempty" json:"TransactionNote"`
      AccountingPeriod *LocalIdType `xml:"AccountingPeriod,omitempty" json:"AccountingPeriod"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    