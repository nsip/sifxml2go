package sifxml


type StaffPersonals []StaffPersonal

    type StaffPersonal struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      LocalId *LocalIdType `xml:"LocalId" json:"LocalId"`
      StateProvinceId *StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      ElectronicIdList *ElectronicIdListType `xml:"ElectronicIdList,omitempty" json:"ElectronicIdList"`
      OtherIdList *OtherIdListType `xml:"OtherIdList,omitempty" json:"OtherIdList"`
      PersonInfo *PersonInfoType `xml:"PersonInfo" json:"PersonInfo"`
      Title *String `xml:"Title,omitempty" json:"Title"`
      EmploymentStatus *AUCodeSetsStaffStatusType `xml:"EmploymentStatus,omitempty" json:"EmploymentStatus"`
      MostRecent *StaffMostRecentContainerType `xml:"MostRecent,omitempty" json:"MostRecent"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    