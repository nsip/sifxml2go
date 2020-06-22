package sifxml


    type StudentAttendanceCollection struct {
        RefId RefIdType `xml:"RefId,attr"`
      StudentAttendanceCollectionYear SchoolYearType `xml:"StudentAttendanceCollectionYear"`
      RoundCode string `xml:"RoundCode"`
      ReportingAuthorityCommonwealthId string `xml:"ReportingAuthorityCommonwealthId"`
      SoftwareVendorInfo SoftwareVendorInfoContainerType `xml:"SoftwareVendorInfo"`
      StudentAttendanceCollectionReportingList StudentAttendanceCollectionReportingListType `xml:"StudentAttendanceCollectionReportingList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    