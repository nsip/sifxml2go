package sifxml


type LEAInfos []LEAInfo

    type LEAInfo struct {
  leainfo `xml:"LEAInfo" json:"LEAInfo"`
}

type leainfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId,omitempty"`
      CommonwealthId *String `xml:"CommonwealthId,omitempty" json:"CommonwealthId,omitempty"`
      LEAName *String `xml:"LEAName" json:"LEAName"`
      LEAURL *String `xml:"LEAURL,omitempty" json:"LEAURL,omitempty"`
      EducationAgencyType *AgencyType `xml:"EducationAgencyType,omitempty" json:"EducationAgencyType,omitempty"`
      LEAContactList *LEAContactListType `xml:"LEAContactList,omitempty" json:"LEAContactList,omitempty"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList,omitempty"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList,omitempty"`
      OperationalStatus *OperationalStatusType `xml:"OperationalStatus,omitempty" json:"OperationalStatus,omitempty"`
      JurisdictionLowerHouse *String `xml:"JurisdictionLowerHouse,omitempty" json:"JurisdictionLowerHouse,omitempty"`
      SLA *AUCodeSetsAustralianStandardGeographicalClassificationASGCType `xml:"SLA,omitempty" json:"SLA,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

