package sifxml


type LibraryPatronStatuss []LibraryPatronStatus

    type LibraryPatronStatus struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LibraryType *String `xml:"LibraryType" json:"LibraryType"`
      PatronRefId *String `xml:"PatronRefId,omitempty" json:"PatronRefId"`
      PatronLocalId *LocalIdType `xml:"PatronLocalId,omitempty" json:"PatronLocalId"`
      PatronRefObject *String `xml:"PatronRefObject" json:"PatronRefObject"`
      PatronName *NameOfRecordType `xml:"PatronName,omitempty" json:"PatronName"`
      ElectronicIdList *ElectronicIdListType `xml:"ElectronicIdList,omitempty" json:"ElectronicIdList"`
      TransactionList *LibraryTransactionListType `xml:"TransactionList,omitempty" json:"TransactionList"`
      MessageList *LibraryMessageListType `xml:"MessageList,omitempty" json:"MessageList"`
      NumberOfCheckouts *Int `xml:"NumberOfCheckouts" json:"NumberOfCheckouts"`
      NumberOfHoldItems *Int `xml:"NumberOfHoldItems,omitempty" json:"NumberOfHoldItems"`
      NumberOfOverdues *Int `xml:"NumberOfOverdues,omitempty" json:"NumberOfOverdues"`
      NumberOfFines *Int `xml:"NumberOfFines,omitempty" json:"NumberOfFines"`
      FineAmount *MonetaryAmountType `xml:"FineAmount,omitempty" json:"FineAmount"`
      NumberOfRefunds *Int `xml:"NumberOfRefunds,omitempty" json:"NumberOfRefunds"`
      RefundAmount *MonetaryAmountType `xml:"RefundAmount,omitempty" json:"RefundAmount"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    