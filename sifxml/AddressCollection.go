package sifxml


type AddressCollections []AddressCollection

    type AddressCollection struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      AddressCollectionYear *SchoolYearType `xml:"AddressCollectionYear" json:"AddressCollectionYear"`
      RoundCode *String `xml:"RoundCode" json:"RoundCode"`
      ReportingAuthorityCommonwealthId *String `xml:"ReportingAuthorityCommonwealthId,omitempty" json:"ReportingAuthorityCommonwealthId"`
      SoftwareVendorInfo *SoftwareVendorInfoContainerType `xml:"SoftwareVendorInfo,omitempty" json:"SoftwareVendorInfo"`
      AddressCollectionReportingList *AddressCollectionReportingListType `xml:"AddressCollectionReportingList,omitempty" json:"AddressCollectionReportingList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    