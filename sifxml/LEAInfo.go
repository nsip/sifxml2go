package sifxml


type LEAInfos []LEAInfo

    type LEAInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      CommonwealthId *String `xml:"CommonwealthId,omitempty" json:"CommonwealthId"`
      LEAName *String `xml:"LEAName" json:"LEAName"`
      LEAURL *String `xml:"LEAURL,omitempty" json:"LEAURL"`
      EducationAgencyType *AgencyType `xml:"EducationAgencyType,omitempty" json:"EducationAgencyType"`
      LEAContactList *LEAContactListType `xml:"LEAContactList,omitempty" json:"LEAContactList"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      OperationalStatus *OperationalStatusType `xml:"OperationalStatus,omitempty" json:"OperationalStatus"`
      JurisdictionLowerHouse *String `xml:"JurisdictionLowerHouse,omitempty" json:"JurisdictionLowerHouse"`
      SLA *AUCodeSetsAustralianStandardGeographicalClassificationASGCType `xml:"SLA,omitempty" json:"SLA"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    