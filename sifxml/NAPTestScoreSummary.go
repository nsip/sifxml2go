package sifxml


type NAPTestScoreSummarys []NAPTestScoreSummary

    type NAPTestScoreSummary struct {
        RefId *RefIdType `xml:"RefId,attr" json:"RefId"`
      SchoolInfoRefId *String `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolACARAId *LocalIdType `xml:"SchoolACARAId" json:"SchoolACARAId"`
      NAPTestRefId *String `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId"`
      NAPTestLocalId *LocalIdType `xml:"NAPTestLocalId" json:"NAPTestLocalId"`
      DomainNationalAverage *Float `xml:"DomainNationalAverage,omitempty" json:"DomainNationalAverage"`
      DomainSchoolAverage *Float `xml:"DomainSchoolAverage,omitempty" json:"DomainSchoolAverage"`
      DomainJurisdictionAverage *Float `xml:"DomainJurisdictionAverage,omitempty" json:"DomainJurisdictionAverage"`
      DomainTopNational60Percent *Float `xml:"DomainTopNational60Percent,omitempty" json:"DomainTopNational60Percent"`
      DomainBottomNational60Percent *Float `xml:"DomainBottomNational60Percent,omitempty" json:"DomainBottomNational60Percent"`
      LocalCodeList *LocalCodeListType `xml:"LocalCodeList,omitempty" json:"LocalCodeList"`
      SIF_Metadata *SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements *SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    