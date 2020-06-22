package sifxml


    type CollectionStatus struct {
        RefId RefIdType `xml:"RefId,attr"`
      ReportingAuthority string `xml:"ReportingAuthority"`
      ReportingAuthoritySystem string `xml:"ReportingAuthoritySystem"`
      ReportingAuthorityCommonwealthId string `xml:"ReportingAuthorityCommonwealthId"`
      SubmittedBy string `xml:"SubmittedBy"`
      SubmissionTimestamp string `xml:"SubmissionTimestamp"`
      AGCollection string `xml:"AGCollection"`
      CollectionYear SchoolYearType `xml:"CollectionYear"`
      RoundCode string `xml:"RoundCode"`
      AGReportingObjectResponseList AGReportingObjectResponseListType `xml:"AGReportingObjectResponseList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    