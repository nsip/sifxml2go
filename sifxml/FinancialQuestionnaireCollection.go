package sifxml


type FinancialQuestionnaireCollections []FinancialQuestionnaireCollection

    type FinancialQuestionnaireCollection struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      FQYear *SchoolYearType `xml:"FQYear,omitempty" json:"FQYear"`
      RoundCode *string `xml:"RoundCode,omitempty" json:"RoundCode"`
      ReportingAuthorityCommonwealthId *string `xml:"ReportingAuthorityCommonwealthId,omitempty" json:"ReportingAuthorityCommonwealthId"`
      SoftwareVendorInfo *SoftwareVendorInfoContainerType `xml:"SoftwareVendorInfo,omitempty" json:"SoftwareVendorInfo"`
      FQReportingList *FQReportingListType `xml:"FQReportingList,omitempty" json:"FQReportingList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    