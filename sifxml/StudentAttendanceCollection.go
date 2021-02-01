package sifxml


type StudentAttendanceCollections []StudentAttendanceCollection

    type StudentAttendanceCollection struct {
  studentattendancecollection `xml:"StudentAttendanceCollection" json:"StudentAttendanceCollection"`
}

type studentattendancecollection struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentAttendanceCollectionYear *SchoolYearType `xml:"StudentAttendanceCollectionYear" json:"StudentAttendanceCollectionYear"`
      RoundCode *String `xml:"RoundCode" json:"RoundCode"`
      SoftwareVendorInfo *SoftwareVendorInfoContainerType `xml:"SoftwareVendorInfo,omitempty" json:"SoftwareVendorInfo,omitempty"`
      StudentAttendanceCollectionReportingList *StudentAttendanceCollectionReportingListType `xml:"StudentAttendanceCollectionReportingList,omitempty" json:"StudentAttendanceCollectionReportingList,omitempty"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList,omitempty"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata,omitempty"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements,omitempty"`
      
      }
    