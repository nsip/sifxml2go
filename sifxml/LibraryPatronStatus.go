package sifxml

import _ "github.com/creasty/defaults"


type LibraryPatronStatuss []LibraryPatronStatus

    type LibraryPatronStatus struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      LibraryType string `xml:"LibraryType" json:"LibraryType"`
      PatronRefId string `xml:"PatronRefId" json:"PatronRefId"`
      PatronLocalId LocalIdType `xml:"PatronLocalId" json:"PatronLocalId"`
      PatronRefObject string `xml:"PatronRefObject" json:"PatronRefObject"`
      PatronName NameOfRecordType `xml:"PatronName,omitempty" json:"PatronName"`
      ElectronicIdList ElectronicIdListType `xml:"ElectronicIdList,omitempty" json:"ElectronicIdList"`
      TransactionList LibraryTransactionListType `xml:"TransactionList,omitempty" json:"TransactionList"`
      MessageList LibraryMessageListType `xml:"MessageList,omitempty" json:"MessageList"`
      NumberOfCheckouts int `xml:"NumberOfCheckouts" json:"NumberOfCheckouts"`
      NumberOfHoldItems int `default:"-2147483648" xml:"NumberOfHoldItems" json:"NumberOfHoldItems"`
      NumberOfOverdues int `default:"-2147483648" xml:"NumberOfOverdues" json:"NumberOfOverdues"`
      NumberOfFines int `default:"-2147483648" xml:"NumberOfFines" json:"NumberOfFines"`
      FineAmount MonetaryAmountType `xml:"FineAmount,omitempty" json:"FineAmount"`
      NumberOfRefunds int `default:"-2147483648" xml:"NumberOfRefunds" json:"NumberOfRefunds"`
      RefundAmount MonetaryAmountType `xml:"RefundAmount,omitempty" json:"RefundAmount"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    