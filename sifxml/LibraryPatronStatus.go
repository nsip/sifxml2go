package sifxml


    type LibraryPatronStatus struct {
        RefId RefIdType `xml:"RefId,attr"`
      LibraryType string `xml:"LibraryType"`
      PatronRefId string `xml:"PatronRefId"`
      PatronLocalId LocalIdType `xml:"PatronLocalId"`
      PatronRefObject string `xml:"PatronRefObject"`
      ElectronicIdList ElectronicIdListType `xml:"ElectronicIdList"`
      TransactionList LibraryTransactionListType `xml:"TransactionList"`
      MessageList LibraryMessageListType `xml:"MessageList"`
      NumberOfCheckouts string `xml:"NumberOfCheckouts"`
      NumberOfOverdues string `xml:"NumberOfOverdues"`
      NumberOfFines string `xml:"NumberOfFines"`
      FineAmount MonetaryAmountType `xml:"FineAmount"`
      NumberOfRefunds string `xml:"NumberOfRefunds"`
      RefundAmount MonetaryAmountType `xml:"RefundAmount"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    