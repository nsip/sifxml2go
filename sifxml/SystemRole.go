package sifxml


type SystemRoles []SystemRole

    type SystemRole struct {
  systemrole `xml:"SystemRole" json:"SystemRole"`
}

type systemrole struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SIF_RefId *SystemRoleObjectRefIdType `xml:"SIF_RefId" json:"SIF_RefId"`
      SystemContextList *SystemContextListType `xml:"SystemContextList" json:"SystemContextList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    