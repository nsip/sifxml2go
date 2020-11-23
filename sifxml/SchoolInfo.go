package sifxml


type SchoolInfos []SchoolInfo

    type SchoolInfo struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      CommonwealthId *String `xml:"CommonwealthId,omitempty" json:"CommonwealthId"`
      ACARAId *String `xml:"ACARAId,omitempty" json:"ACARAId"`
      OtherIdList *OtherIdListType `xml:"OtherIdList,omitempty" json:"OtherIdList"`
      SchoolName *String `xml:"SchoolName" json:"SchoolName"`
      LEAInfoRefId *RefIdType `xml:"LEAInfoRefId,omitempty" json:"LEAInfoRefId"`
      OtherLEA *SchoolInfo_OtherLEA
      SchoolDistrict *String `xml:"SchoolDistrict,omitempty" json:"SchoolDistrict"`
      SchoolDistrictLocalId *LocalIdType `xml:"SchoolDistrictLocalId,omitempty" json:"SchoolDistrictLocalId"`
      SchoolType *AUCodeSetsSchoolLevelType `xml:"SchoolType,omitempty" json:"SchoolType"`
      SchoolFocusList *SchoolFocusListType `xml:"SchoolFocusList,omitempty" json:"SchoolFocusList"`
      SchoolURL *SchoolURLType `xml:"SchoolURL,omitempty" json:"SchoolURL"`
      SchoolEmailList *EmailListType `xml:"SchoolEmailList,omitempty" json:"SchoolEmailList"`
      PrincipalInfo *PrincipalInfoType `xml:"PrincipalInfo,omitempty" json:"PrincipalInfo"`
      SchoolContactList *SchoolContactListType `xml:"SchoolContactList,omitempty" json:"SchoolContactList"`
      AddressList *AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      PhoneNumberList *PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      SessionType *AUCodeSetsSessionTypeType `xml:"SessionType,omitempty" json:"SessionType"`
      YearLevels *YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      ARIA *Float `xml:"ARIA,omitempty" json:"ARIA"`
      OperationalStatus *OperationalStatusType `xml:"OperationalStatus,omitempty" json:"OperationalStatus"`
      FederalElectorate *AUCodeSetsFederalElectorateType `xml:"FederalElectorate,omitempty" json:"FederalElectorate"`
      Campus *CampusContainerType `xml:"Campus,omitempty" json:"Campus"`
      SchoolSector *AUCodeSetsSchoolSectorCodeType `xml:"SchoolSector" json:"SchoolSector"`
      IndependentSchool *AUCodeSetsYesOrNoCategoryType `xml:"IndependentSchool,omitempty" json:"IndependentSchool"`
      NonGovSystemicStatus *AUCodeSetsSystemicStatusType `xml:"NonGovSystemicStatus,omitempty" json:"NonGovSystemicStatus"`
      System *AUCodeSetsSchoolSystemType `xml:"System,omitempty" json:"System"`
      ReligiousAffiliation *AUCodeSetsAustralianStandardClassificationOfReligiousGroupsASCRGType `xml:"ReligiousAffiliation,omitempty" json:"ReligiousAffiliation"`
      SchoolGeographicLocation *AUCodeSetsSchoolLocationType `xml:"SchoolGeographicLocation,omitempty" json:"SchoolGeographicLocation"`
      LocalGovernmentArea *String `xml:"LocalGovernmentArea,omitempty" json:"LocalGovernmentArea"`
      JurisdictionLowerHouse *String `xml:"JurisdictionLowerHouse,omitempty" json:"JurisdictionLowerHouse"`
      SLA *AUCodeSetsAustralianStandardGeographicalClassificationASGCType `xml:"SLA,omitempty" json:"SLA"`
      SchoolCoEdStatus *AUCodeSetsSchoolCoEdStatusType `xml:"SchoolCoEdStatus,omitempty" json:"SchoolCoEdStatus"`
      BoardingSchoolStatus *AUCodeSetsYesOrNoCategoryType `xml:"BoardingSchoolStatus,omitempty" json:"BoardingSchoolStatus"`
      YearLevelEnrollmentList *YearLevelEnrollmentListType `xml:"YearLevelEnrollmentList,omitempty" json:"YearLevelEnrollmentList"`
      TotalEnrollments *TotalEnrollmentsType `xml:"TotalEnrollments,omitempty" json:"TotalEnrollments"`
      Entity_Open *String `xml:"Entity_Open,omitempty" json:"Entity_Open"`
      Entity_Close *String `xml:"Entity_Close,omitempty" json:"Entity_Close"`
      SchoolGroupList *SchoolGroupListType `xml:"SchoolGroupList,omitempty" json:"SchoolGroupList"`
      SchoolTimeZone *AUCodeSetsAustralianTimeZoneType `xml:"SchoolTimeZone,omitempty" json:"SchoolTimeZone"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type SchoolInfo_OtherLEA struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *RefIdType `xml:",chardata" json:"value"`
}
