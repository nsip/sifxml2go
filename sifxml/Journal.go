package sifxml


type Journals []Journal

    type Journal struct {
  journal `xml:"Journal" json:"Journal"`
}

type journal struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      DebitFinancialAccountRefId *String `xml:"DebitFinancialAccountRefId,omitempty" json:"DebitFinancialAccountRefId,omitempty"`
      CreditFinancialAccountRefId *String `xml:"CreditFinancialAccountRefId,omitempty" json:"CreditFinancialAccountRefId,omitempty"`
      DebitAccountCode *String `xml:"DebitAccountCode,omitempty" json:"DebitAccountCode,omitempty"`
      CreditAccountCode *String `xml:"CreditAccountCode,omitempty" json:"CreditAccountCode,omitempty"`
      OriginatingTransactionRefId *Journal_OriginatingTransactionRefId
      Amount *MonetaryAmountType `xml:"Amount" json:"Amount"`
      GSTCodeOriginal *String `xml:"GSTCodeOriginal,omitempty" json:"GSTCodeOriginal,omitempty"`
      GSTCodeReplacement *String `xml:"GSTCodeReplacement,omitempty" json:"GSTCodeReplacement,omitempty"`
      Note *String `xml:"Note,omitempty" json:"Note,omitempty"`
      JournalAdjustmentList *JournalAdjustmentListType `xml:"JournalAdjustmentList,omitempty" json:"JournalAdjustmentList,omitempty"`
      CreatedDate *String `xml:"CreatedDate,omitempty" json:"CreatedDate,omitempty"`
      ApprovedDate *String `xml:"ApprovedDate,omitempty" json:"ApprovedDate,omitempty"`
      CreatedBy *String `xml:"CreatedBy,omitempty" json:"CreatedBy,omitempty"`
      ApprovedBy *String `xml:"ApprovedBy,omitempty" json:"ApprovedBy,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    type Journal_OriginatingTransactionRefId struct {
  journal_originatingtransactionrefid `xml:"Journal_OriginatingTransactionRefId" json:"Journal_OriginatingTransactionRefId"`
}

type journal_originatingtransactionrefid struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
