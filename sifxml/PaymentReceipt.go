package sifxml


    type PaymentReceipt struct {
        RefId RefIdType `xml:"RefId,attr"`
      LocalId LocalIdType `xml:"LocalId"`
      TransactionType string `xml:"TransactionType"`
      InvoiceRefId string `xml:"InvoiceRefId"`
      PaymentReceiptLineList PaymentReceiptLineListType `xml:"PaymentReceiptLineList"`
      VendorInfoRefId string `xml:"VendorInfoRefId"`
      DebtorRefId string `xml:"DebtorRefId"`
      ChargedLocationInfoRefId string `xml:"ChargedLocationInfoRefId"`
      TransactionDate string `xml:"TransactionDate"`
      TransactionAmount DebitOrCreditAmountType `xml:"TransactionAmount"`
      ReceivedTransactionId string `xml:"ReceivedTransactionId"`
      FinancialAccountRefIdList FinancialAccountRefIdListType `xml:"FinancialAccountRefIdList"`
      AccountCodeList AccountCodeListType `xml:"AccountCodeList"`
      TransactionDescription string `xml:"TransactionDescription"`
      TaxRate string `xml:"TaxRate"`
      TaxAmount MonetaryAmountType `xml:"TaxAmount"`
      TransactionMethod string `xml:"TransactionMethod"`
      ChequeNumber string `xml:"ChequeNumber"`
      TransactionNote string `xml:"TransactionNote"`
      AccountingPeriod LocalIdType `xml:"AccountingPeriod"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    