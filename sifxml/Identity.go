package sifxml


type Identitys []Identity

    type Identity struct {
  identity `xml:"Identity" json:"Identity"`
}

type identity struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SIF_RefId *Identity_SIF_RefId `xml:"SIF_RefId" json:"SIF_RefId"`
      AuthenticationSource *String `xml:"AuthenticationSource" json:"AuthenticationSource"`
      IdentityAssertions *IdentityAssertionsType `xml:"IdentityAssertions,omitempty" json:"IdentityAssertions,omitempty"`
      PasswordList *PasswordListType `xml:"PasswordList,omitempty" json:"PasswordList,omitempty"`
      AuthenticationSourceGlobalUID *String `xml:"AuthenticationSourceGlobalUID,omitempty" json:"AuthenticationSourceGlobalUID,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    
type Identity_SIF_RefId struct {
  identity_sif_refid `xml:"Identity_SIF_RefId" json:"Identity_SIF_RefId"`
}

type identity_sif_refid struct {
      SIF_RefObject *String `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value *String `xml:",chardata" json:"value"`
}

