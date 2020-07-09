package sifxml


type LibraryPatronStatuss []LibraryPatronStatus

    type LibraryPatronStatus struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LibraryType *string `xml:"LibraryType,omitempty" json:"LibraryType"`
      PatronRefId *string `xml:"PatronRefId,omitempty" json:"PatronRefId"`
      PatronLocalId *LocalIdType `xml:"PatronLocalId,omitempty" json:"PatronLocalId"`
      PatronRefObject *string `xml:"PatronRefObject,omitempty" json:"PatronRefObject"`
      ElectronicIdList *ElectronicIdListType `xml:"ElectronicIdList,omitempty" json:"ElectronicIdList"`
      TransactionList *LibraryTransactionListType `xml:"TransactionList,omitempty" json:"TransactionList"`
      MessageList *LibraryMessageListType `xml:"MessageList,omitempty" json:"MessageList"`
      NumberOfCheckouts *int `xml:"NumberOfCheckouts,omitempty" json:"NumberOfCheckouts"`
      NumberOfOverdues *int `xml:"NumberOfOverdues,omitempty" json:"NumberOfOverdues"`
      NumberOfFines *int `xml:"NumberOfFines,omitempty" json:"NumberOfFines"`
      FineAmount *MonetaryAmountType `xml:"FineAmount,omitempty" json:"FineAmount"`
      NumberOfRefunds *int `xml:"NumberOfRefunds,omitempty" json:"NumberOfRefunds"`
      RefundAmount *MonetaryAmountType `xml:"RefundAmount,omitempty" json:"RefundAmount"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    