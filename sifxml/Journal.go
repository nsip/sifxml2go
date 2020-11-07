package sifxml


type Journals []Journal

    type Journal struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      DebitFinancialAccountRefId *string `xml:"DebitFinancialAccountRefId,omitempty" json:"DebitFinancialAccountRefId"`
      CreditFinancialAccountRefId *string `xml:"CreditFinancialAccountRefId,omitempty" json:"CreditFinancialAccountRefId"`
      DebitAccountCode *string `xml:"DebitAccountCode,omitempty" json:"DebitAccountCode"`
      CreditAccountCode *string `xml:"CreditAccountCode,omitempty" json:"CreditAccountCode"`
      OriginatingTransactionRefId *Journal_OriginatingTransactionRefId
      Amount *MonetaryAmountType `xml:"Amount" json:"Amount"`
      GSTCodeOriginal *string `xml:"GSTCodeOriginal,omitempty" json:"GSTCodeOriginal"`
      GSTCodeReplacement *string `xml:"GSTCodeReplacement,omitempty" json:"GSTCodeReplacement"`
      Note *string `xml:"Note,omitempty" json:"Note"`
      JournalAdjustmentList *JournalAdjustmentListType `xml:"JournalAdjustmentList,omitempty" json:"JournalAdjustmentList"`
      CreatedDate *string `xml:"CreatedDate,omitempty" json:"CreatedDate"`
      ApprovedDate *string `xml:"ApprovedDate,omitempty" json:"ApprovedDate"`
      CreatedBy *string `xml:"CreatedBy,omitempty" json:"CreatedBy"`
      ApprovedBy *string `xml:"ApprovedBy,omitempty" json:"ApprovedBy"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type Journal_OriginatingTransactionRefId struct {
      SIF_RefObject *string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *string `xml:",chardata" json:"value"`
}
