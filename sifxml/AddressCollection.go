package sifxml


    type AddressCollection struct {
        RefId RefIdType `xml:"RefId,attr"`
      AddressCollectionYear SchoolYearType `xml:"AddressCollectionYear"`
      RoundCode string `xml:"RoundCode"`
      ReportingAuthorityCommonwealthId string `xml:"ReportingAuthorityCommonwealthId"`
      SoftwareVendorInfo SoftwareVendorInfoContainerType `xml:"SoftwareVendorInfo"`
      AddressCollectionReportingList AddressCollectionReportingListType `xml:"AddressCollectionReportingList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    