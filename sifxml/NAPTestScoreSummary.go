package sifxml


    type NAPTestScoreSummary struct {
        RefId RefIdType `xml:"RefId,attr"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId"`
      SchoolACARAId LocalIdType `xml:"SchoolACARAId"`
      NAPTestRefId string `xml:"NAPTestRefId"`
      NAPTestLocalId LocalIdType `xml:"NAPTestLocalId"`
      DomainNationalAverage string `xml:"DomainNationalAverage"`
      DomainSchoolAverage string `xml:"DomainSchoolAverage"`
      DomainJurisdictionAverage string `xml:"DomainJurisdictionAverage"`
      DomainTopNational60Percent string `xml:"DomainTopNational60Percent"`
      DomainBottomNational60Percent string `xml:"DomainBottomNational60Percent"`
      LocalCodeList LocalCodeListType `xml:"LocalCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements"`
      
      }
    