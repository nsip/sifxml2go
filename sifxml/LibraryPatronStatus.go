package sifxml


type LibraryPatronStatuss []LibraryPatronStatus

    type LibraryPatronStatus struct {
  librarypatronstatus `xml:"LibraryPatronStatus"`
}

type librarypatronstatus struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LibraryType *String `xml:"LibraryType" json:"LibraryType"`
      PatronRefId *String `xml:"PatronRefId,omitempty" json:"PatronRefId,omitempty"`
      PatronLocalId *LocalIdType `xml:"PatronLocalId,omitempty" json:"PatronLocalId,omitempty"`
      PatronRefObject *String `xml:"PatronRefObject" json:"PatronRefObject"`
      PatronName *NameOfRecordType `xml:"PatronName,omitempty" json:"PatronName,omitempty"`
      ElectronicIdList *ElectronicIdListType `xml:"ElectronicIdList,omitempty" json:"ElectronicIdList,omitempty"`
      TransactionList *LibraryTransactionListType `xml:"TransactionList,omitempty" json:"TransactionList,omitempty"`
      MessageList *LibraryMessageListType `xml:"MessageList,omitempty" json:"MessageList,omitempty"`
      NumberOfCheckouts *Int `xml:"NumberOfCheckouts" json:"NumberOfCheckouts"`
      NumberOfHoldItems *Int `xml:"NumberOfHoldItems,omitempty" json:"NumberOfHoldItems,omitempty"`
      NumberOfOverdues *Int `xml:"NumberOfOverdues,omitempty" json:"NumberOfOverdues,omitempty"`
      NumberOfFines *Int `xml:"NumberOfFines,omitempty" json:"NumberOfFines,omitempty"`
      FineAmount *MonetaryAmountType `xml:"FineAmount,omitempty" json:"FineAmount,omitempty"`
      NumberOfRefunds *Int `xml:"NumberOfRefunds,omitempty" json:"NumberOfRefunds,omitempty"`
      RefundAmount *MonetaryAmountType `xml:"RefundAmount,omitempty" json:"RefundAmount,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

