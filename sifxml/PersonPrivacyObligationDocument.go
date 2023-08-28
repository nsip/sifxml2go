package sifxml


type PersonPrivacyObligationDocuments []PersonPrivacyObligationDocument

    type PersonPrivacyObligationDocument struct {
  personprivacyobligationdocument `xml:"PersonPrivacyObligationDocument"`
}

type personprivacyobligationdocument struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      ParentRefId *String `xml:"ParentRefId" json:"ParentRefId"`
      ParentObjectTypeName *String `xml:"ParentObjectTypeName" json:"ParentObjectTypeName"`
      SchoolYear *SchoolYearType `xml:"SchoolYear" json:"SchoolYear"`
      StartDate *String `xml:"StartDate,omitempty" json:"StartDate,omitempty"`
      EndDate *String `xml:"EndDate,omitempty" json:"EndDate,omitempty"`
      SettingLocationList *SettingLocationListType `xml:"SettingLocationList,omitempty" json:"SettingLocationList,omitempty"`
      ContactForRequestsRefId *String `xml:"ContactForRequestsRefId" json:"ContactForRequestsRefId"`
      ContactForRequestsObjectTypeName *String `xml:"ContactForRequestsObjectTypeName" json:"ContactForRequestsObjectTypeName"`
      ConsentToSharingOfData *ConsentToSharingOfDataContainerType `xml:"ConsentToSharingOfData" json:"ConsentToSharingOfData"`
      PermissionToParticipateList *PermissionToParticipateListType `xml:"PermissionToParticipateList" json:"PermissionToParticipateList"`
      ApplicableLawList *ApplicableLawListType `xml:"ApplicableLawList,omitempty" json:"ApplicableLawList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

