package sifxml


type SchoolInfos []SchoolInfo

    type SchoolInfo struct {
  schoolinfo `xml:"SchoolInfo" json:"SchoolInfo"`
}

type schoolinfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId,omitempty"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId,omitempty"`
      CommonwealthId *String `xml:"CommonwealthId,omitempty" json:"CommonwealthId,omitempty"`
      ParentCommonwealthId *String `xml:"ParentCommonwealthId,omitempty" json:"ParentCommonwealthId,omitempty"`
      ACARAId *String `xml:"ACARAId,omitempty" json:"ACARAId,omitempty"`
      OtherIdList *OtherIdListType `xml:"OtherIdList,omitempty" json:"OtherIdList,omitempty"`
      SchoolName *String `xml:"SchoolName" json:"SchoolName"`
      LEAInfoRefId *RefIdType `xml:"LEAInfoRefId,omitempty" json:"LEAInfoRefId,omitempty"`
      OtherLEA *SchoolInfo_OtherLEA
      SchoolDistrict *String `xml:"SchoolDistrict,omitempty" json:"SchoolDistrict,omitempty"`
      SchoolDistrictLocalId *LocalIdType `xml:"SchoolDistrictLocalId,omitempty" json:"SchoolDistrictLocalId,omitempty"`
      SchoolType *AUCodeSetsSchoolLevelType `xml:"SchoolType,omitempty" json:"SchoolType,omitempty"`
      SchoolFocusList *SchoolFocusListType `xml:"SchoolFocusList,omitempty" json:"SchoolFocusList,omitempty"`
      SchoolURL *SchoolURLType `xml:"SchoolURL,omitempty" json:"SchoolURL,omitempty"`
      SchoolEmailList *EmailListType `xml:"SchoolEmailList,omitempty" json:"SchoolEmailList,omitempty"`
      PrincipalInfo *PrincipalInfoType `xml:"PrincipalInfo,omitempty" json:"PrincipalInfo,omitempty"`
      SchoolContactList *SchoolContactListType `xml:"SchoolContactList,omitempty" json:"SchoolContactList,omitempty"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList,omitempty"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList,omitempty"`
      SessionType *AUCodeSetsSessionTypeType `xml:"SessionType,omitempty" json:"SessionType,omitempty"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels,omitempty"`
      ARIA *Float `xml:"ARIA,omitempty" json:"ARIA,omitempty"`
      OperationalStatus *OperationalStatusType `xml:"OperationalStatus,omitempty" json:"OperationalStatus,omitempty"`
      FederalElectorate *AUCodeSetsFederalElectorateType `xml:"FederalElectorate,omitempty" json:"FederalElectorate,omitempty"`
      Campus *CampusContainerType `xml:"Campus,omitempty" json:"Campus,omitempty"`
      SchoolSector *AUCodeSetsSchoolSectorCodeType `xml:"SchoolSector" json:"SchoolSector"`
      IndependentSchool *AUCodeSetsYesOrNoCategoryType `xml:"IndependentSchool,omitempty" json:"IndependentSchool,omitempty"`
      NonGovSystemicStatus *AUCodeSetsSystemicStatusType `xml:"NonGovSystemicStatus,omitempty" json:"NonGovSystemicStatus,omitempty"`
      System *AUCodeSetsSchoolSystemType `xml:"System,omitempty" json:"System,omitempty"`
      ReligiousAffiliation *AUCodeSetsAustralianStandardClassificationOfReligiousGroupsASCRGType `xml:"ReligiousAffiliation,omitempty" json:"ReligiousAffiliation,omitempty"`
      SchoolGeographicLocation *AUCodeSetsSchoolLocationType `xml:"SchoolGeographicLocation,omitempty" json:"SchoolGeographicLocation,omitempty"`
      LocalGovernmentArea *String `xml:"LocalGovernmentArea,omitempty" json:"LocalGovernmentArea,omitempty"`
      JurisdictionLowerHouse *String `xml:"JurisdictionLowerHouse,omitempty" json:"JurisdictionLowerHouse,omitempty"`
      SLA *AUCodeSetsAustralianStandardGeographicalClassificationASGCType `xml:"SLA,omitempty" json:"SLA,omitempty"`
      SchoolCoEdStatus *AUCodeSetsSchoolCoEdStatusType `xml:"SchoolCoEdStatus,omitempty" json:"SchoolCoEdStatus,omitempty"`
      BoardingSchoolStatus *AUCodeSetsYesOrNoCategoryType `xml:"BoardingSchoolStatus,omitempty" json:"BoardingSchoolStatus,omitempty"`
      YearLevelEnrollmentList *YearLevelEnrollmentListType `xml:"YearLevelEnrollmentList,omitempty" json:"YearLevelEnrollmentList,omitempty"`
      TotalEnrollments *TotalEnrollmentsType `xml:"TotalEnrollments,omitempty" json:"TotalEnrollments,omitempty"`
      Entity_Open *String `xml:"Entity_Open,omitempty" json:"Entity_Open,omitempty"`
      Entity_Close *String `xml:"Entity_Close,omitempty" json:"Entity_Close,omitempty"`
      SchoolGroupList *SchoolGroupListType `xml:"SchoolGroupList,omitempty" json:"SchoolGroupList,omitempty"`
      SchoolTimeZone *AUCodeSetsAustralianTimeZoneType `xml:"SchoolTimeZone,omitempty" json:"SchoolTimeZone,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    
type SchoolInfo_OtherLEA struct {
  schoolinfo_otherlea `xml:"SchoolInfo_OtherLEA" json:"SchoolInfo_OtherLEA"`
}

type schoolinfo_otherlea struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *RefIdType `xml:",chardata" json:"value"`
}

