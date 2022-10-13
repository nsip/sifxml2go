package sifxml


type AddressCollections []AddressCollection

    type AddressCollection struct {
  addresscollection `xml:"AddressCollection" json:"AddressCollection"`
}

type addresscollection struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      AddressCollectionYear *SchoolYearType `xml:"AddressCollectionYear" json:"AddressCollectionYear"`
      RoundCode *String `xml:"RoundCode" json:"RoundCode"`
      SoftwareVendorInfo *SoftwareVendorInfoContainerType `xml:"SoftwareVendorInfo,omitempty" json:"SoftwareVendorInfo,omitempty"`
      AddressCollectionReportingList *AddressCollectionReportingListType `xml:"AddressCollectionReportingList,omitempty" json:"AddressCollectionReportingList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

