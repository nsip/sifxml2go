package sifxml


type CollectionDeclarations []CollectionDeclaration

    type CollectionDeclaration struct {
  collectiondeclaration `xml:"CollectionDeclaration"`
}

type collectiondeclaration struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      ReportingAuthorityList *ReportingAuthorityListType `xml:"ReportingAuthorityList" json:"ReportingAuthorityList"`
      SubmittingAuthority *ReportingAuthorityType `xml:"SubmittingAuthority,omitempty" json:"SubmittingAuthority,omitempty"`
      SubmittedBy *SignatoryType `xml:"SubmittedBy,omitempty" json:"SubmittedBy,omitempty"`
      Recipient *String `xml:"Recipient,omitempty" json:"Recipient,omitempty"`
      Collection *String `xml:"Collection,omitempty" json:"Collection,omitempty"`
      CollectionYear *SchoolYearType `xml:"CollectionYear" json:"CollectionYear"`
      RoundCode *String `xml:"RoundCode,omitempty" json:"RoundCode,omitempty"`
      Declaration *String `xml:"Declaration" json:"Declaration"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

