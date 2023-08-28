package sifxml


type CollectionStatuss []CollectionStatus

    type CollectionStatus struct {
  collectionstatus `xml:"CollectionStatus"`
}

type collectionstatus struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      ReportingAuthority *String `xml:"ReportingAuthority" json:"ReportingAuthority"`
      ReportingAuthoritySystem *String `xml:"ReportingAuthoritySystem,omitempty" json:"ReportingAuthoritySystem,omitempty"`
      ReportingAuthorityCommonwealthId *String `xml:"ReportingAuthorityCommonwealthId,omitempty" json:"ReportingAuthorityCommonwealthId,omitempty"`
      SubmittedBy *String `xml:"SubmittedBy,omitempty" json:"SubmittedBy,omitempty"`
      SubmissionTimestamp *String `xml:"SubmissionTimestamp,omitempty" json:"SubmissionTimestamp,omitempty"`
      AGCollection *String `xml:"AGCollection,omitempty" json:"AGCollection,omitempty"`
      CollectionYear *SchoolYearType `xml:"CollectionYear" json:"CollectionYear"`
      RoundCode *String `xml:"RoundCode" json:"RoundCode"`
      AGReportingObjectResponseList *AGReportingObjectResponseListType `xml:"AGReportingObjectResponseList,omitempty" json:"AGReportingObjectResponseList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    

