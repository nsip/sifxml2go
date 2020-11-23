package sifxml


type CensusCollections []CensusCollection

    type CensusCollection struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      CensusYear *SchoolYearType `xml:"CensusYear" json:"CensusYear"`
      RoundCode *String `xml:"RoundCode" json:"RoundCode"`
      ReportingAuthorityCommonwealthId *String `xml:"ReportingAuthorityCommonwealthId,omitempty" json:"ReportingAuthorityCommonwealthId"`
      SoftwareVendorInfo *SoftwareVendorInfoContainerType `xml:"SoftwareVendorInfo,omitempty" json:"SoftwareVendorInfo"`
      CensusReportingList *CensusReportingListType `xml:"CensusReportingList,omitempty" json:"CensusReportingList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    