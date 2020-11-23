package sifxml


type CollectionStatuss []CollectionStatus

    type CollectionStatus struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      ReportingAuthority *String `xml:"ReportingAuthority" json:"ReportingAuthority"`
      ReportingAuthoritySystem *String `xml:"ReportingAuthoritySystem,omitempty" json:"ReportingAuthoritySystem"`
      ReportingAuthorityCommonwealthId *String `xml:"ReportingAuthorityCommonwealthId,omitempty" json:"ReportingAuthorityCommonwealthId"`
      SubmittedBy *String `xml:"SubmittedBy,omitempty" json:"SubmittedBy"`
      SubmissionTimestamp *String `xml:"SubmissionTimestamp,omitempty" json:"SubmissionTimestamp"`
      AGCollection *String `xml:"AGCollection,omitempty" json:"AGCollection"`
      CollectionYear *SchoolYearType `xml:"CollectionYear" json:"CollectionYear"`
      RoundCode *String `xml:"RoundCode" json:"RoundCode"`
      AGReportingObjectResponseList *AGReportingObjectResponseListType `xml:"AGReportingObjectResponseList,omitempty" json:"AGReportingObjectResponseList"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    