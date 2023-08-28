package sifxml


type CollectionAcquittals []CollectionAcquittal

    type CollectionAcquittal struct {
  collectionacquittal `xml:"CollectionAcquittal"`
}

type collectionacquittal struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      ReportingAuthorityList *ReportingAuthorityListType `xml:"ReportingAuthorityList" json:"ReportingAuthorityList"`
      SubmittingAuthority *ReportingAuthorityType `xml:"SubmittingAuthority,omitempty" json:"SubmittingAuthority,omitempty"`
      SubmittedBy *SignatoryType `xml:"SubmittedBy,omitempty" json:"SubmittedBy,omitempty"`
      AuditedBy *SignatoryType `xml:"AuditedBy,omitempty" json:"AuditedBy,omitempty"`
      AuditorASICNumber *String `xml:"AuditorASICNumber,omitempty" json:"AuditorASICNumber,omitempty"`
      Recipient *String `xml:"Recipient,omitempty" json:"Recipient,omitempty"`
      Collection *String `xml:"Collection,omitempty" json:"Collection,omitempty"`
      CollectionYear *SchoolYearType `xml:"CollectionYear" json:"CollectionYear"`
      RoundCode *String `xml:"RoundCode,omitempty" json:"RoundCode,omitempty"`
      Declaration *String `xml:"Declaration" json:"Declaration"`
      AuditorStatement *String `xml:"AuditorStatement" json:"AuditorStatement"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

