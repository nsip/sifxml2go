package sifxml


type SystemRoles []SystemRole

    type SystemRole struct {
  systemrole `xml:"SystemRole" json:"SystemRole"`
}

type systemrole struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SIF_RefId *SystemRole_SIF_RefId `xml:"SIF_RefId" json:"SIF_RefId"`
      SystemContextList *SystemRole_SystemContextList `xml:"SystemContextList" json:"SystemContextList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    type SystemRole_SIF_RefId struct {
  systemrole_sif_refid `xml:"SystemRole_SIF_RefId" json:"SystemRole_SIF_RefId"`
}

type systemrole_sif_refid struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
type SystemRole_SystemContextList struct {
  systemrole_systemcontextlist `xml:"SystemRole_SystemContextList" json:"SystemRole_SystemContextList"`
}

type systemrole_systemcontextlist struct {
      SystemContext []SystemRole_SystemContext `xml:"SystemContext" json:"SystemContext"`
}
type SystemRole_SystemContext struct {
  systemrole_systemcontext `xml:"SystemRole_SystemContext" json:"SystemRole_SystemContext"`
}

type systemrole_systemcontext struct {
      SystemId *String `xml:"SystemId,attr" json:"SystemId"`
      RoleList *SystemRole_RoleList
}
type SystemRole_RoleList struct {
  systemrole_rolelist `xml:"SystemRole_RoleList" json:"SystemRole_RoleList"`
}

type systemrole_rolelist struct {
      Role []SystemRole_Role `xml:"Role" json:"Role"`
}
type SystemRole_Role struct {
  systemrole_role `xml:"SystemRole_Role" json:"SystemRole_Role"`
}

type systemrole_role struct {
      RoleId *String `xml:"RoleId,attr" json:"RoleId"`
      RoleScopeList *SystemRole_RoleScopeList
}
type SystemRole_RoleScopeList struct {
  systemrole_rolescopelist `xml:"SystemRole_RoleScopeList" json:"SystemRole_RoleScopeList"`
}

type systemrole_rolescopelist struct {
      RoleScope []SystemRole_RoleScope `xml:"RoleScope" json:"RoleScope"`
}
type SystemRole_RoleScope struct {
  systemrole_rolescope `xml:"SystemRole_RoleScope" json:"SystemRole_RoleScope"`
}

type systemrole_rolescope struct {
       RoleScopeName *String `xml:"RoleScopeName,omitempty" json:"RoleScopeName,omitempty"`
      RoleScopeRefId *SystemRole_RoleScopeRefId
}
type SystemRole_RoleScopeRefId struct {
  systemrole_rolescoperefid `xml:"SystemRole_RoleScopeRefId" json:"SystemRole_RoleScopeRefId"`
}

type systemrole_rolescoperefid struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}
