package sifxml

import _ "github.com/creasty/defaults"


type Identitys []Identity

    type Identity struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      SIF_RefId Identity_SIF_RefId `xml:"SIF_RefId" json:"SIF_RefId"`
      AuthenticationSource string `xml:"AuthenticationSource" json:"AuthenticationSource"`
      IdentityAssertions IdentityAssertionsType `xml:"IdentityAssertions,omitempty" json:"IdentityAssertions"`
      PasswordList PasswordListType `xml:"PasswordList,omitempty" json:"PasswordList"`
      AuthenticationSourceGlobalUID string `xml:"AuthenticationSourceGlobalUID,omitempty" json:"AuthenticationSourceGlobalUID"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type Identity_SIF_RefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"SIF_RefObject"`
      Value string `xml:",chardata" json:"value"`
}
