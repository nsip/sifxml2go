package sifxml


type SystemRoles []SystemRole

    type SystemRole struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SIF_RefId *SystemRole_SIF_RefId `xml:"SIF_RefId" json:"SIF_RefId"`
      SystemContextList *SystemRole_SystemContextList `xml:"SystemContextList" json:"SystemContextList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type SystemRole_SIF_RefId struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
type SystemRole_SystemContextList struct {
      SystemContext []SystemRole_SystemContext `xml:"SystemContext" json:"SystemContext"`
}
type SystemRole_SystemContext struct {
      SystemId *String `xml:"SystemId,attr" json:"SystemId"`
      RoleList *SystemRole_RoleList
}
type SystemRole_RoleList struct {
      Role []SystemRole_Role `xml:"Role" json:"Role"`
}
type SystemRole_Role struct {
      RoleId *String `xml:"RoleId,attr" json:"RoleId"`
      RoleScopeList *SystemRole_RoleScopeList
}
type SystemRole_RoleScopeList struct {
      RoleScope []SystemRole_RoleScope `xml:"RoleScope" json:"RoleScope"`
}
type SystemRole_RoleScope struct {
       RoleScopeName *String `xml:"RoleScopeName,omitempty" json:"RoleScopeName"`
      RoleScopeRefId *SystemRole_RoleScopeRefId
}
type SystemRole_RoleScopeRefId struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
