package sifxml


type NAPTestScoreSummarys []NAPTestScoreSummary

    type NAPTestScoreSummary struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolACARAId *LocalIdType `xml:"SchoolACARAId" json:"SchoolACARAId"`
      NAPTestRefId *string `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId"`
      NAPTestLocalId *LocalIdType `xml:"NAPTestLocalId" json:"NAPTestLocalId"`
      DomainNationalAverage *float64 `xml:"DomainNationalAverage" json:"DomainNationalAverage"`
      DomainSchoolAverage *float64 `xml:"DomainSchoolAverage" json:"DomainSchoolAverage"`
      DomainJurisdictionAverage *float64 `xml:"DomainJurisdictionAverage" json:"DomainJurisdictionAverage"`
      DomainTopNational60Percent *float64 `xml:"DomainTopNational60Percent" json:"DomainTopNational60Percent"`
      DomainBottomNational60Percent *float64 `xml:"DomainBottomNational60Percent" json:"DomainBottomNational60Percent"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    