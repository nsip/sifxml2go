package sifxml


type FinancialQuestionnaireCollections []FinancialQuestionnaireCollection

    type FinancialQuestionnaireCollection struct {
  financialquestionnairecollection `xml:"FinancialQuestionnaireCollection" json:"FinancialQuestionnaireCollection"`
}

type financialquestionnairecollection struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      FQYear *SchoolYearType `xml:"FQYear" json:"FQYear"`
      RoundCode *String `xml:"RoundCode" json:"RoundCode"`
      SoftwareVendorInfo *SoftwareVendorInfoContainerType `xml:"SoftwareVendorInfo" json:"SoftwareVendorInfo"`
      FQReportingList *FQReportingListType `xml:"FQReportingList,omitempty" json:"FQReportingList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

