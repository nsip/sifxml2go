package sifxml


type ChargedLocationInfos []ChargedLocationInfo

    type ChargedLocationInfo struct {
  chargedlocationinfo `xml:"ChargedLocationInfo" json:"ChargedLocationInfo"`
}

type chargedlocationinfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocationType *String `xml:"LocationType" json:"LocationType"`
      SiteCategory *String `xml:"SiteCategory" json:"SiteCategory"`
      Name *String `xml:"Name" json:"Name"`
      Description *String `xml:"Description,omitempty" json:"Description,omitempty"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId,omitempty"`
      ParentChargedLocationInfoRefId *String `xml:"ParentChargedLocationInfoRefId,omitempty" json:"ParentChargedLocationInfoRefId,omitempty"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId,omitempty"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList,omitempty"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

