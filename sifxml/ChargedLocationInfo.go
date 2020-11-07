package sifxml


type ChargedLocationInfos []ChargedLocationInfo

    type ChargedLocationInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocationType *string `xml:"LocationType" json:"LocationType"`
      SiteCategory *string `xml:"SiteCategory" json:"SiteCategory"`
      Name *string `xml:"Name" json:"Name"`
      Description *string `xml:"Description,omitempty" json:"Description"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      ParentChargedLocationInfoRefId *string `xml:"ParentChargedLocationInfoRefId,omitempty" json:"ParentChargedLocationInfoRefId"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    