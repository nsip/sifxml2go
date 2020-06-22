package sifxml


    type FinancialQuestionnaireCollection struct {
        RefId RefIdType `xml:"RefId,attr"`
      FQYear SchoolYearType `xml:"FQYear"`
      RoundCode string `xml:"RoundCode"`
      ReportingAuthorityCommonwealthId string `xml:"ReportingAuthorityCommonwealthId"`
      SoftwareVendorInfo SoftwareVendorInfoContainerType `xml:"SoftwareVendorInfo"`
      FQReportingList FQReportingListType `xml:"FQReportingList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    