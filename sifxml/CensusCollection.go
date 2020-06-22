package sifxml


    type CensusCollection struct {
        RefId RefIdType `xml:"RefId,attr"`
      CensusYear SchoolYearType `xml:"CensusYear"`
      RoundCode string `xml:"RoundCode"`
      ReportingAuthorityCommonwealthId string `xml:"ReportingAuthorityCommonwealthId"`
      SoftwareVendorInfo SoftwareVendorInfoContainerType `xml:"SoftwareVendorInfo"`
      CensusReportingList CensusReportingListType `xml:"CensusReportingList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    