package sifxml


type ChargedLocationInfos []ChargedLocationInfo

    type ChargedLocationInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocationType *String `xml:"LocationType" json:"LocationType"`
      SiteCategory *String `xml:"SiteCategory" json:"SiteCategory"`
      Name *String `xml:"Name" json:"Name"`
      Description *String `xml:"Description,omitempty" json:"Description"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      ParentChargedLocationInfoRefId *String `xml:"ParentChargedLocationInfoRefId,omitempty" json:"ParentChargedLocationInfoRefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    