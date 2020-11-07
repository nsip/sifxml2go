package sifxml

import _ "github.com/creasty/defaults"


type AddressCollections []AddressCollection

    type AddressCollection struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      AddressCollectionYear SchoolYearType `xml:"AddressCollectionYear" json:"AddressCollectionYear"`
      RoundCode string `xml:"RoundCode" json:"RoundCode"`
      ReportingAuthorityCommonwealthId string `xml:"ReportingAuthorityCommonwealthId,omitempty" json:"ReportingAuthorityCommonwealthId"`
      SoftwareVendorInfo SoftwareVendorInfoContainerType `xml:"SoftwareVendorInfo,omitempty" json:"SoftwareVendorInfo"`
      AddressCollectionReportingList AddressCollectionReportingListType `xml:"AddressCollectionReportingList,omitempty" json:"AddressCollectionReportingList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    