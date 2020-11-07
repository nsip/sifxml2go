package sifxml

import _ "github.com/creasty/defaults"


type StudentAttendanceCollections []StudentAttendanceCollection

    type StudentAttendanceCollection struct {
        RefId RefIdType `xml:"RefId,attr" json:"RefId"`
      StudentAttendanceCollectionYear SchoolYearType `xml:"StudentAttendanceCollectionYear" json:"StudentAttendanceCollectionYear"`
      RoundCode string `xml:"RoundCode" json:"RoundCode"`
      ReportingAuthorityCommonwealthId string `xml:"ReportingAuthorityCommonwealthId,omitempty" json:"ReportingAuthorityCommonwealthId"`
      SoftwareVendorInfo SoftwareVendorInfoContainerType `xml:"SoftwareVendorInfo,omitempty" json:"SoftwareVendorInfo"`
      StudentAttendanceCollectionReportingList StudentAttendanceCollectionReportingListType `xml:"StudentAttendanceCollectionReportingList,omitempty" json:"StudentAttendanceCollectionReportingList"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    