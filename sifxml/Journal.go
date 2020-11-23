package sifxml


type Journals []Journal

    type Journal struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      DebitFinancialAccountRefId *String `xml:"DebitFinancialAccountRefId,omitempty" json:"DebitFinancialAccountRefId"`
      CreditFinancialAccountRefId *String `xml:"CreditFinancialAccountRefId,omitempty" json:"CreditFinancialAccountRefId"`
      DebitAccountCode *String `xml:"DebitAccountCode,omitempty" json:"DebitAccountCode"`
      CreditAccountCode *String `xml:"CreditAccountCode,omitempty" json:"CreditAccountCode"`
      OriginatingTransactionRefId *Journal_OriginatingTransactionRefId
      Amount *MonetaryAmountType `xml:"Amount" json:"Amount"`
      GSTCodeOriginal *String `xml:"GSTCodeOriginal,omitempty" json:"GSTCodeOriginal"`
      GSTCodeReplacement *String `xml:"GSTCodeReplacement,omitempty" json:"GSTCodeReplacement"`
      Note *String `xml:"Note,omitempty" json:"Note"`
      JournalAdjustmentList *JournalAdjustmentListType `xml:"JournalAdjustmentList,omitempty" json:"JournalAdjustmentList"`
      CreatedDate *String `xml:"CreatedDate,omitempty" json:"CreatedDate"`
      ApprovedDate *String `xml:"ApprovedDate,omitempty" json:"ApprovedDate"`
      CreatedBy *String `xml:"CreatedBy,omitempty" json:"CreatedBy"`
      ApprovedBy *String `xml:"ApprovedBy,omitempty" json:"ApprovedBy"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type Journal_OriginatingTransactionRefId struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
