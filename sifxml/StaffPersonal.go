package sifxml


type StaffPersonals []StaffPersonal

    type StaffPersonal struct {
  staffpersonal `xml:"StaffPersonal" json:"StaffPersonal"`
}

type staffpersonal struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId,omitempty"`
      ElectronicIdList *ElectronicIdListType `xml:"ElectronicIdList,omitempty" json:"ElectronicIdList,omitempty"`
      OtherIdList *OtherIdListType `xml:"OtherIdList,omitempty" json:"OtherIdList,omitempty"`
      PersonInfo *PersonInfoType `xml:"PersonInfo" json:"PersonInfo"`
      Title *String `xml:"Title,omitempty" json:"Title,omitempty"`
      EmploymentStatus *AUCodeSetsStaffStatusType `xml:"EmploymentStatus,omitempty" json:"EmploymentStatus,omitempty"`
      MostRecent *StaffMostRecentContainerType `xml:"MostRecent,omitempty" json:"MostRecent,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

