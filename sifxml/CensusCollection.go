package sifxml


type CensusCollections []CensusCollection

    type CensusCollection struct {
  censuscollection `xml:"CensusCollection"`
}

type censuscollection struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      CensusYear *SchoolYearType `xml:"CensusYear" json:"CensusYear"`
      RoundCode *String `xml:"RoundCode" json:"RoundCode"`
      SoftwareVendorInfo *SoftwareVendorInfoContainerType `xml:"SoftwareVendorInfo,omitempty" json:"SoftwareVendorInfo,omitempty"`
      CensusReportingList *CensusReportingListType `xml:"CensusReportingList,omitempty" json:"CensusReportingList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

