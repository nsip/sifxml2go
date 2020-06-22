package sifxml


    type Journal struct {
        RefId RefIdType `xml:"RefId,attr"`
      LocalId LocalIdType `xml:"LocalId"`
      DebitFinancialAccountRefId string `xml:"DebitFinancialAccountRefId"`
      CreditFinancialAccountRefId string `xml:"CreditFinancialAccountRefId"`
      DebitAccountCode string `xml:"DebitAccountCode"`
      CreditAccountCode string `xml:"CreditAccountCode"`
      OriginatingTransactionRefId Journal_OriginatingTransactionRefId `xml:"OriginatingTransactionRefId"`
      Amount MonetaryAmountType `xml:"Amount"`
      GSTCodeOriginal string `xml:"GSTCodeOriginal"`
      GSTCodeReplacement string `xml:"GSTCodeReplacement"`
      Note string `xml:"Note"`
      JournalAdjustmentList JournalAdjustmentListType `xml:"JournalAdjustmentList"`
      CreatedDate string `xml:"CreatedDate"`
      ApprovedDate string `xml:"ApprovedDate"`
      CreatedBy string `xml:"CreatedBy"`
      ApprovedBy string `xml:"ApprovedBy"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    type Journal_OriginatingTransactionRefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr"`
      Value string `xml:",chardata"`
}
